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
      "debian-stretch-*",
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
    "aws-marketplace",
  ]
}

resource "random_id" "ethstat_secret" {
  byte_length = 16
}

resource "aws_sqs_queue" "faketime_queue" {
  name                        = "faketime-${random_id.bucket_postfix.hex}"
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

EOF

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common jq",
      "sudo curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -",
      "sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable\"",
      "sudo apt-get update",
      "sudo apt-get install -y docker-ce docker-ce-cli containerd.io jq python3-pip",
      "sudo pip3 install awscli --upgrade > /dev/null",
      "sudo systemctl start docker",
      "sudo gpasswd -a admin docker",
      "printf 'FROM alpine\nCOPY --from=trajano/alpine-libfaketime  /faketime.so /lib/faketime.so\n' > /tmp/Dockerfile.libfaketime",
      "sudo docker build -f /tmp/Dockerfile.libfaketime . -t libfaketime:latest",
      "sudo docker run -v $PWD:/tmp --rm --entrypoint cp libfaketime:latest /lib/faketime.so /tmp/libfaketime.so",
      "sudo aws --region ${var.region} s3 cp libfaketime.so s3://${local.bastion_bucket}/libs/libfaketime.so",
      "for value in ${join(" ", var.faketime)}; do aws --region ${var.region} sqs send-message --queue-url ${aws_sqs_queue.faketime_queue.id} --message-body \\\"$value\\\" --message-attributes \"{ \\\"faketimeValue\\\":{ \\\"DataType\\\":\\\"String\\\",\\\"StringValue\\\":\\\"$value\\\"}}\" ; done",
    ]

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "admin"
      private_key = "${tls_private_key.ssh.private_key_pem}"
      timeout     = "10m"
    }
  }


  tags = "${merge(local.common_tags, map("Name", local.default_bastion_resource_name))}"
}

resource "random_string" "random" {
  length = 16
  special = false
}

resource "local_file" "bootstrap" {
  filename = "${path.module}/generated-bootstrap.sh"

  content = <<EOF
#!/bin/bash

set -e

# Kill all running containers
if [ $(docker ps -a -q | wc -l ) -gt 0 ]; then
  sudo docker stop $(docker ps -a -q)
  sudo docker rm $(docker ps -a -q)
fi

sudo curl -Ls "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
echo "Pull docker images ..."
sudo docker pull ${local.quorum_docker_image} > /dev/null
sudo docker pull ${local.ethstats_docker_image}> /dev/null
sudo docker pull prom/prometheus > /dev/null
sudo docker pull prom/node-exporter:latest > /dev/null
sudo docker pull grafana/grafana:latest > /dev/null
sudo docker pull hunterlong/gethexporter:latest > /dev/null
echo "Done"
sudo mkdir -p /opt/prometheus
sudo mkdir -p /opt/grafana/dashboards
sudo mkdir -p /opt/grafana/provisioning/dashboards
sudo mkdir -p /opt/grafana/provisioning/datasources
sudo curl -Ls https://grafana.com/api/dashboards/6976/revisions/3/download -o /opt/grafana/dashboards/dashboard-geth.json
sudo curl -Ls https://grafana.com/api/dashboards/1860/revisions/14/download -o /opt/grafana/dashboards/dashboard-node-exporter.json
sudo docker run -d -e "WS_SECRET=${random_id.ethstat_secret.hex}" -p ${local.ethstats_port}:${local.ethstats_port} ${local.ethstats_docker_image}

export AWS_DEFAULT_REGION=${var.region}
export TASK_REVISION=${aws_ecs_task_definition.quorum.revision}
sudo rm -rf ${local.shared_volume_container_path}
sudo mkdir -p ${local.shared_volume_container_path}/mappings
sudo mkdir -p ${local.privacy_addresses_folder}

count=0
while [ $count -lt ${var.number_of_nodes} ]
do
  count=$(ls ${local.privacy_addresses_folder} | grep ^ip | wc -l)
  sudo aws --region ${var.region} s3 cp --recursive s3://${local.s3_revision_folder}/ ${local.shared_volume_container_path}/ > /dev/null 2>&1 \
    | echo Wait for nodes in Quorum network being up ... $count/${var.number_of_nodes}
  sleep 1
done

if which jq >/dev/null; then
  echo "Found jq"
else
  echo "jq not found. Instaling ..."
  sudo apt-get -y install jq
fi

count=30
inc_num=0
while [ $count -gt $inc_num ]
do
  status=$(aws --region ${var.region} ecs describe-clusters --clusters ${local.ecs_cluster_name} | jq -r .clusters[].status)
  if [ $status == "ACTIVE" ]; then
    inc_num=$count
  fi
  sleep 1
  inc_num=$((inc_num+1))
done

for t in $(aws --region ${var.region} ecs list-tasks --cluster ${local.ecs_cluster_name} | jq -r .taskArns[])
do
  task_metadata=$(aws --region ${var.region} ecs describe-tasks --cluster ${local.ecs_cluster_name} --tasks $t)
  # works with awsvpc mode
  # HOST_IP=$(echo $task_metadata | jq -r '.tasks[0] | .containers[] | select(.name == "${local.quorum_run_container_name}") | .networkInterfaces[] | .privateIpv4Address')
  if [ "${var.ecs_mode}" == "EC2" ]
  then
    CONTAINER_INSTANCE_ARN=$(aws --region ${var.region} ecs describe-tasks --tasks $t --cluster ${local.ecs_cluster_name} | jq -r '.tasks[] | .containerInstanceArn')
    EC2_INSTANCE_ID=$(aws --region ${var.region} ecs  describe-container-instances --container-instances $CONTAINER_INSTANCE_ARN --cluster ${local.ecs_cluster_name} |jq -r '.containerInstances[] | .ec2InstanceId')
    HOST_IP=$(aws --region ${var.region} ec2 describe-instances --instance-ids $EC2_INSTANCE_ID | jq -r '.Reservations[0] | .Instances[] | .PublicIpAddress')
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
- job_name: gethexporter
  static_configs:
  - targets:
    - gethexporter:9090
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
    gethexporter:
        image: hunterlong/gethexporter
        environment:
            - GETH=http://$ip:${local.quorum_rpc_port}
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

  depends_on = ["aws_ecs_cluster.quorum"]

  provisioner "remote-exec" {
    script = "${local_file.bootstrap.filename}"

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "admin"
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
