locals {
  host_ip_file         = "${local.shared_volume_container_path}/host_ip"
  task_revision_file   = "${local.shared_volume_container_path}/task_revision"
  service_file         = "${local.shared_volume_container_path}/service"
  account_address_file = "${local.shared_volume_container_path}/first_account_address"
  hosts_folder         = "${local.shared_volume_container_path}/hosts"
  libfaketime_folder   = "${local.shared_volume_container_path}/lib"
  libfaketime_file     = "${local.shared_volume_container_path}/lib/libfaketime_value"
  node_info_folder     = "${local.shared_volume_container_path}/nodeinfo"
  parity_docker_hb_config_generator = "${var.parity_docker_hb_config_generator}"

  metadata_bootstrap_container_status_file = "${local.shared_volume_container_path}/metadata_bootstrap_container_status"

  // For S3 related operations
  s3_revision_folder = "${local.parity_bucket}/rev_$TASK_REVISION"
  s3_libfaketime_file = "${local.bastion_bucket}/libs/libfaketime.so"
  normalized_host_ip = "ip_$(echo $HOST_IP | sed -e 's/\\./_/g')"

  node_key_bootstrap_commands = [
    "mkdir -p ${local.shared_volume_container_path}/prepare",
    "cd ${local.shared_volume_container_path}/prepare",
    "/bin/hbbft_config_generator ${var.number_of_nodes} Docker",
  ]

  node_key_bootstrap_container_definition = {
    name      = "${local.node_key_bootstrap_container_name}"
    image     = "${local.parity_docker_hb_config_generator}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"
 
        options = {
        awslogs-group         = "${aws_cloudwatch_log_group.parity.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
     }
   }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    environments = []

    portMappings = []

    volumesFrom = []

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -f ${local.node_id_file} ];",
      ]
    }

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.node_key_bootstrap_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 512
  }

  metadata_bootstrap_commands = [
    "set -e",
    "apk update",
    "apk add curl jq",
    "while [ ! -f \"${local.shared_volume_container_path}/prepare/reserved-peers\" ]; do sleep 1; done",
    "export TASK_REVISION=",
    "echo exporting TASK_REVISION ....",
    "while [ -z \"$TASK_REVISION\" ]; do TASK_REVISION=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq .Revision -r); sleep 1; done",
    "echo ECS_CONTAINER_METADATA_URI: $ECS_CONTAINER_METADATA_URI/task",
    "echo \"Task Revision: $TASK_REVISION\"",
    "echo $TASK_REVISION > ${local.task_revision_file}",
    "export HOST_IP=$(/usr/bin/curl http://169.254.169.254/latest/meta-data/public-ipv4)",
    "echo \"Host IP: $HOST_IP\"",
    "echo $HOST_IP > ${local.host_ip_file}",
    "export TASK_ARN=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')",
    "export REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')",
    "export SERVICE_GROUP=$(aws ecs describe-tasks --region ${var.region} --cluster ${local.ecs_cluster_name} --tasks $TASK_ARN | jq -r '.tasks[0] | .group')",
    "echo $SERVICE_GROUP > ${local.service_file}",
    "mkdir -p ${local.hosts_folder}",
    "mkdir -p ${local.config_folder}",
    "mkdir -p ${local.node_info_folder}",
    "mkdir -p ${local.keys_folder}",
    "mkdir -p ${local.libfaketime_folder}",

    "mkdir -p ${local.parity_data_dir}/network/",
    "mkdir -p ${local.parity_data_dir}/keys/${local.chain_name}",

    "count=0; while [ $count -lt 1 ]; do count=$(ls ${local.libfaketime_folder} | grep libfaketime.so | wc -l); aws s3 cp s3://${local.s3_libfaketime_file} ${local.libfaketime_folder}/libfaketime.so > /dev/null 2>&1 | echo \"Wait for libfaketime to appear on S3 ... \"; sleep 1; done",
    "touch ${local.libfaketime_file}",
    "export CLOCK_SKEW=$(aws sqs --region ${var.region} receive-message --queue-url ${aws_sqs_queue.faketime_queue.id} --visibility-timeout=300 | jq .Messages[].Body | tr -d '\\\"')",
    "echo $CLOCK_SKEW > ${local.libfaketime_file}",
    "echo \"$(echo $SERVICE_GROUP | sed 's/.*://')  ip=$HOST_IP  clock_skew=$CLOCK_SKEW chaos_testing_command=${join(" ", var.chaos_testing_run_command)}\" > ${local.node_info_folder}/${local.normalized_host_ip}",

    "aws s3 cp ${local.node_info_folder}/${local.normalized_host_ip} s3://${local.s3_revision_folder}/nodeinfo/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "aws s3 cp ${local.host_ip_file} s3://${local.s3_revision_folder}/hosts/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
   
    // Gather all IPs
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.hosts_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/hosts ${local.hosts_folder} > /dev/null 2>&1 | echo \"Wait for other containers to report their IPs ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All containers have reported their IPs\"",

    //check if bootnode is first
    "firt_host_ip=$(ls ${local.hosts_folder} | grep ^ip | sort | head -1)",
    "if [ $firt_host_ip == ${local.normalized_host_ip} ]; then i=0; for f in $(ls ${local.hosts_folder} | grep ^ip | sort); do i=$((i+1)); ip=$(cat ${local.hosts_folder}/$f); sed -i s\"/127.0.0.1:3030$i/$ip:30303/g\" ${local.shared_volume_container_path}/prepare/hbbft_validator_$i.toml; sed -i s\"/^port = 85.*/port = 8545/\" ${local.shared_volume_container_path}/prepare/hbbft_validator_$i.toml;sed -i \"0,/disable/s/disable = true/disable = false/\" ${local.shared_volume_container_path}/prepare/hbbft_validator_$i.toml; sed -i s\"/^port = 303.*/port = 30303/\" ${local.shared_volume_container_path}/prepare/hbbft_validator_$i.toml; sed -i s\"/127.0.0.1:3030$i/$ip:30303/\" ${local.shared_volume_container_path}/prepare/reserved-peers; aws s3 cp ${local.shared_volume_container_path}/prepare/hbbft_validator_$i.toml s3://${local.s3_revision_folder}/config/$f --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}; aws s3 cp ${local.shared_volume_container_path}/prepare/hbbft_validator_key_$i s3://${local.s3_revision_folder}/keys/$f --sse aws:kms --sse-kms-key-id  ${aws_kms_key.bucket.arn}; aws s3 cp ${local.shared_volume_container_path}/prepare/hbbft_validator_key_$i.json s3://${local.s3_revision_folder}/keys_json/$f --sse aws:kms --sse-kms-key-id  ${aws_kms_key.bucket.arn} ; aws s3 cp ${local.shared_volume_container_path}/prepare/reserved-peers s3://${local.s3_revision_folder}/config/reserved-peers --sse aws:kms --sse-kms-key-id  ${aws_kms_key.bucket.arn}; done; fi",

    // Gather config
    "count=0; while [ ! -f ${local.config_file} ]; do aws s3 cp s3://${local.s3_revision_folder}/config/${local.normalized_host_ip} ${local.config_file}  | echo \"Wait for prepare toml config ... \" ; sleep 1; done",
    "count=0; while [ ! -f ${local.shared_volume_container_path}/reserved-peers ]; do aws s3 cp s3://${local.s3_revision_folder}/config/reserved-peers  ${local.shared_volume_container_path} | echo \"Wait for prepare config reserved-peers ... \"; sleep 1; done",
    "echo \"All nodes have registered configs\"",

    // Gather key
    "count=0; while [ ! -f ${local.parity_data_dir}/network/key ]; do aws s3 cp s3://${local.s3_revision_folder}/keys/${local.normalized_host_ip} ${local.parity_data_dir}/network/key > /dev/null 2>&1 | echo \"Wait download key ...\"; sleep 1; done",
    "count=0; while [ ! -f ${local.parity_data_dir}/keys/${local.chain_name}/hbbft_validator_key.json ]; do aws s3 cp s3://${local.s3_revision_folder}/keys_json/${local.normalized_host_ip} ${local.parity_data_dir}/keys/${local.chain_name}/hbbft_validator_key.json > /dev/null 2>&1 | echo \"Wait for download hbbft_validator_key.json ...\"; sleep 1; done",

    "echo \"All nodes have registered their keys\"",

    // Prepare Genesis file

    //"echo '${replace(jsonencode(local.genesis), "/\"(true|false|[0-9]+)\"/", "$1")}' > ${local.genesis_file}",
    //"sed -i s'/RANDOM_NETWORK_ID/${random_integer.network_id.result}/' ${local.genesis_file}",
    
    "curl http://gostomski.pl/spec.json -o ${local.genesis_file}",
    
    //"count=0; while [ -f ${local.genesis_file} ]; do aws s3 cp s3://${local.s3_revision_folder}/config/spec.json ${local.genesis_file} | echo \"Download spec.json \"; done",
    "sed -i s/DPoSChain/${local.chain_name}/ ${local.genesis_file}",
    //"sed -i s/difficulty-value/${var.genesis_difficulty}/ ${local.genesis_file}",
    //"sed -i s/gasLimit-value/${var.genesis_gas_limit}/ ${local.genesis_file}",
    //"sed -i s/timestamp-value/${var.genesis_timestamp}/ ${local.genesis_file}",
    //"sed -i s/networkID-value/${random_integer.network_id.result}/ ${local.genesis_file}",
   
    "echo ${local.genesis_file}",
    "cat ${local.genesis_file}",

    "aws s3 cp ${local.shared_volume_container_path}/prepare/hbbft_validator_key_1.json s3://${local.s3_revision_folder}/privacyaddresses/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",

    // Write status
    "echo \"Done!\" > ${local.metadata_bootstrap_container_status_file}",
    "echo Done",
  ]

  metadata_bootstrap_container_definition = {
    name      = "${local.metadata_bootstrap_container_name}"
    image     = "${local.aws_cli_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"
 
        options = {
         awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
         awslogs-region        = "${var.region}"
         awslogs-stream-prefix = "logs"
      }
    }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    environments = []

    portMappings = []

    volumesFrom = [
      {
        sourceContainer = "${local.node_key_bootstrap_container_name}"
      },
    ]

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -f ${local.metadata_bootstrap_container_status_file} ];",
      ]
    }

    entryPoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.metadata_bootstrap_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 512
  }
}
