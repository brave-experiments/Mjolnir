locals {
  default_bastion_resource_name = "${format("quorum-bastion-%s", var.network_name)}"
  ethstats_docker_image         = "puppeth/ethstats:latest"
  ethstats_port                 = 3000
  bastion_bucket    = "${var.region}-bastion-${lower(var.network_name)}-${random_id.bucket_postfix.hex}"
}

data "aws_ami" "this" {
  most_recent = true

  filter {
    name = "name"

    values = [
      "amzn2-ami-hvm-*",
    ]
  }

  filter {
    name = "virtualization-type"

    values = [
      "hvm",
    ]
  }

  filter {
    name = "architecture"

    values = [
      "x86_64",
    ]
  }

  owners = [
    "137112412989",
  ]

  # amazon
}

resource "random_id" "ethstat_secret" {
  byte_length = 16
}

resource "tls_private_key" "ssh" {
  algorithm = "RSA"
  rsa_bits  = "2048"
}

resource "aws_key_pair" "ssh" {
  public_key = "${tls_private_key.ssh.public_key_openssh}"
  key_name   = "${local.default_bastion_resource_name}"
}

resource "local_file" "private_key" {
  filename = "${path.module}/quorum-${var.network_name}.pem"
  content  = "${tls_private_key.ssh.private_key_pem}"

  provisioner "local-exec" {
    on_failure = "continue"
    command    = "chmod 600 ${self.filename}"
  }
}

resource "aws_instance" "bastion" {
  ami           = "${data.aws_ami.this.id}"
  instance_type = "t2.large"

  vpc_security_group_ids = [
    "${aws_security_group.quorum.id}",
    "${aws_security_group.bastion-ssh.id}",
    "${aws_security_group.bastion-ethstats.id}",
  ]

  subnet_id                   = "${aws_subnet.public.id}"
  associate_public_ip_address = "true"
  key_name                    = "${aws_key_pair.ssh.key_name}"
  iam_instance_profile        = "${aws_iam_instance_profile.bastion.name}"

  user_data = <<EOF
#!/bin/bash

set -e

# START: added per suggestion from AWS support to mitigate an intermittent failures from yum update
sleep 20
yum clean all
yum repolist
# END

yum -y update
yum -y install jq
amazon-linux-extras install docker -y
systemctl enable docker
systemctl start docker

curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker pull ${local.quorum_docker_image}
docker pull prom/prometheus
docker pull prom/node-exporter:latest
docker pull grafana/grafana:latest
docker pull ethereum/client-go:latest
mkdir -p /opt/prometheus
mkdir -p /opt/grafana/dashboards
mkdir -p /opt/grafana/provisioning/dashboards
mkdir -p /opt/grafana/provisioning/datasources
curl -L https://gist.githubusercontent.com/karalabe/e7ca79abdec54755ceae09c08bd090cd/raw/3a400ab90f9402f2233280afd086cb9d6aac2111/dashboard.json -o /opt/grafana/dashboards/dashboard-geth.json
curl -L https://grafana.com/api/dashboards/1860/revisions/14/download -o /opt/grafana/dashboards/dashboard-node-exporter.json
docker run -d -e "WS_SECRET=${random_id.ethstat_secret.hex}" -p ${local.ethstats_port}:${local.ethstats_port} ${local.ethstats_docker_image}
EOF

  provisioner "remote-exec" {
    inline = [
      "sudo yum -y update",
      "sudo yum -y install jq",
      "sudo amazon-linux-extras install docker -y",
      "sudo systemctl enable docker",
      "sudo systemctl start docker",
      "printf 'FROM alpine\nCOPY --from=trajano/alpine-libfaketime  /faketime.so /lib/faketime.so\n' > /tmp/Dockerfile.libfaketime",
      "sudo docker build -f /tmp/Dockerfile.libfaketime . -t libfaketime:latest",
      "sudo docker run -v $PWD:/tmp --rm --entrypoint cp libfaketime:latest /lib/faketime.so /tmp/libfaketime.so",
      "sudo aws s3 cp libfaketime.so s3://${local.bastion_bucket}/libs/libfaketime.so"
    ]

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "ec2-user"
      private_key = "${tls_private_key.ssh.private_key_pem}"
      timeout     = "10m"
    }
  }


  tags = "${merge(local.common_tags, map("Name", local.default_bastion_resource_name))}"
}

resource "random_string" "random" {
  length = 16
  special = true
}

resource "local_file" "bootstrap" {
  filename = "${path.module}/generated-bootstrap.sh"

  content = <<EOF
#!/bin/bash

set -e

export AWS_DEFAULT_REGION=${var.region}
export TASK_REVISION=${aws_ecs_task_definition.quorum.revision}
sudo rm -rf ${local.shared_volume_container_path}
sudo mkdir -p ${local.shared_volume_container_path}/mappings
sudo mkdir -p ${local.privacy_addresses_folder}

# Faketime array ( ClockSkew )
old_IFS=$IFS
IFS=',' faketime=(${join(" ", var.faketime)})
IFS=$${old_IFS}
counter="$${#faketime[@]}"

while [ $counter -gt 0 ]
do
    echo -n "$${faketime[-1]}" > ./$counter
    faketime=($${faketime[@]::$counter})
    sudo aws s3 cp ./$counter s3://${local.bastion_bucket}/clockSkew/
    counter=$((counter - 1))
done

count=0
while [ $count -lt ${var.number_of_nodes} ]
do
  count=$(ls ${local.privacy_addresses_folder} | grep ^ip | wc -l)
  sudo aws s3 cp --recursive s3://${local.s3_revision_folder}/ ${local.shared_volume_container_path}/ > /dev/null 2>&1 \
    | echo Wait for nodes in Quorum network being up ... $count/${var.number_of_nodes}
  sleep 1
done

if which jq >/dev/null; then
  echo "Found jq"
else
  echo "jq not found. Instaling ..."
  sudo yum -y install jq
fi

for t in $(aws ecs list-tasks --cluster ${local.ecs_cluster_name} | jq -r .taskArns[])
do
  task_metadata=$(aws ecs describe-tasks --cluster ${local.ecs_cluster_name} --tasks $t)
  HOST_IP=$(echo $task_metadata | jq -r '.tasks[0] | .containers[] | select(.name == "${local.quorum_run_container_name}") | .networkInterfaces[] | .privateIpv4Address')
  if [ "${var.ecs_mode}" == "EC2" ]
  then
    CONTAINER_INSTANCE_ARN=$(aws ecs describe-tasks --tasks $t --cluster ${local.ecs_cluster_name} | jq -r '.tasks[] | .containerInstanceArn')
    EC2_INSTANCE_ID=$(aws ecs  describe-container-instances --container-instances $CONTAINER_INSTANCE_ARN --cluster ${local.ecs_cluster_name} |jq -r '.containerInstances[] | .ec2InstanceId')
    HOST_IP=$(aws ec2 describe-instances --instance-ids $EC2_INSTANCE_ID | jq -r '.Reservations[0] | .Instances[] | .PublicIpAddress')
  fi
  group=$(echo $task_metadata | jq -r '.tasks[0] | .group')
  taskArn=$(echo $task_metadata | jq -r '.tasks[0] | .taskDefinitionArn')
  # only care about new task
  if [[ "$taskArn" == *:$TASK_REVISION ]]; then
     echo $group | sudo tee ${local.shared_volume_container_path}/mappings/${local.normalized_host_ip}
  fi
done

cat <<SS | sudo tee ${local.shared_volume_container_path}/quorum_metadata
quorum:
  nodes:
SS
nodes=(${join(" ", aws_ecs_service.quorum.*.name)})
cd ${local.shared_volume_container_path}/mappings
for idx in "$${!nodes[@]}"
do
  f=$(grep -l $${nodes[$idx]} *)
  ip=$(cat ${local.hosts_folder}/$f)
  nodeIdx=$((idx+1))
  script="/usr/local/bin/Node$nodeIdx"
  cat <<SS | sudo tee $script
#!/bin/bash

sudo docker run --rm -it ${local.quorum_docker_image} attach http://$ip:${local.quorum_rpc_port} $@
SS
  sudo chmod +x $script
  cat <<SS | sudo tee -a ${local.shared_volume_container_path}/quorum_metadata
    Node$nodeIdx:
      privacy-address: $(cat ${local.privacy_addresses_folder}/$f)
      url: http://$ip:${local.quorum_rpc_port}
      third-party-url: http://$ip:${local.tessera_thirdparty_port}
SS
done

cat <<SS | sudo tee /opt/prometheus/prometheus.yml
global:
  scrape_interval:     15s # By default, scrape targets every 15 seconds.

  # Attach these labels to any time series or alerts when communicating with
  # external systems (federation, remote storage, Alertmanager).
  external_labels:
    monitor: 'monitor'

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
- job_name: geth
  metrics_path: /debug/metrics/prometheus
  scheme: http
  static_configs:
  - targets:
    - geth:6060
- job_name: 'node'
  static_configs:
  - targets: [ node-exporter:9100 ]
  file_sd_configs:
  - files:
    - 'targets.json'
SS

cat <<SS | sudo tee /opt/prometheus/docker-compose.yml
# docker-compose.yml
version: '2'
services:
    prometheus:
        image: prom/prometheus:latest
        volumes:
            - /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
            - /opt/prometheus/targets.json:/etc/prometheus/targets.json
        command:
            - '--config.file=/etc/prometheus/prometheus.yml'
        ports:
            - '9090:9090'
    node-exporter:
        image: prom/node-exporter:latest
        ports:
            - '9100:9100'
    grafana:
        image: grafana/grafana:latest
        volumes:
            - /opt/grafana/dashboards:/var/lib/grafana/dashboards
            - /opt/grafana/provisioning/dashboards/all.yml:/etc/grafana/provisioning/dashboards/all.yml
            - /opt/grafana/provisioning/datasources/all.yml:/etc/grafana/provisioning/datasources/all.yml
        environment:
            - GF_SECURITY_ADMIN_PASSWORD=${random_string.random.result}
        depends_on:
            - prometheus
        ports:
            - '3001:3000'
    geth:
        image: ethereum/client-go:latest
        ports:
            - '6060:6060'
        command: --goerli --metrics --metrics.expensive --pprof --pprofaddr=0.0.0.0

SS

count=$(ls ${local.privacy_addresses_folder} | grep ^ip | wc -l)
target_file=/tmp/targets.json
i=0
echo '[' > $target_file
for idx in "$${!nodes[@]}"
do
  f=$(grep -l $${nodes[$idx]} *)
  ip=$(cat ${local.hosts_folder}/$f)
  i=$(($i+1))
  if [ $i -lt "$count" ]; then
    echo '{ "targets": ["'$ip':9100"] },' >> $target_file
  else
    echo '{ "targets": ["'$ip':9100"] }'  >> $target_file
  fi
done
echo ']' >> $target_file
sudo mv $target_file /opt/prometheus/

cat <<SS | sudo tee /opt/grafana/provisioning/datasources/all.yml
datasources:
- name: 'prometheus'
  type: 'prometheus'
  access: 'proxy'
  org_id: 1
  url: 'http://prometheus:9090'
  is_default: true
  version: 1
  editable: true
SS

cat <<SS | sudo tee /opt/grafana/provisioning/dashboards/all.yml
- name: 'default'
  org_id: 1
  folder: ''
  type: 'file'
  options:
    folder: '/var/lib/grafana/dashboards'
SS

sudo sed -i s'/datasource":.*/datasource" :"prometheus",/' /opt/grafana/dashboards/dashboard-geth.json
sudo sed -i s'/datasource":.*/datasource" :"prometheus",/' /opt/grafana/dashboards/dashboard-node-exporter.json
sudo /usr/local/bin/docker-compose -f /opt/prometheus/docker-compose.yml up -d --force-recreate
EOF
}

resource "null_resource" "bastion_remote_exec" {
  triggers {
    bastion             = "${aws_instance.bastion.public_dns}"
    ecs_task_definition = "${aws_ecs_task_definition.quorum.revision}"
    script              = "${md5(local_file.bootstrap.content)}"
  }

  provisioner "remote-exec" {
    script = "${local_file.bootstrap.filename}"

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "ec2-user"
      private_key = "${tls_private_key.ssh.private_key_pem}"
      timeout     = "10m"
    }
  }
}

resource "aws_s3_bucket" "bastion" {
  bucket        = "${local.bastion_bucket}"
  region        = "${var.region}"
  force_destroy = true

  versioning {
    enabled = true
  }
}
