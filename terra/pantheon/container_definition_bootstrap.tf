locals {
  host_ip_file         = "${local.shared_volume_container_path}/host_ip"
  task_revision_file   = "${local.shared_volume_container_path}/task_revision"
  service_file         = "${local.shared_volume_container_path}/service"
  account_address_file = "${local.shared_volume_container_path}/first_account_address"
  hosts_folder         = "${local.shared_volume_container_path}/hosts"
  libfaketime_folder   = "${local.shared_volume_container_path}/lib"
  libfaketime_file     = "${local.shared_volume_container_path}/lib/libfaketime_value"

  metadata_bootstrap_container_status_file = "${local.shared_volume_container_path}/metadata_bootstrap_container_status"

  // For S3 related operations
  s3_revision_folder = "${local.pantheon_bucket}/rev_$TASK_REVISION"
  s3_libfaketime_file = "${local.bastion_bucket}/libs/libfaketime.so"
  normalized_host_ip = "ip_$(echo $HOST_IP | sed -e 's/\\./_/g')"

  node_key_bootstrap_commands = [
    "mkdir -p ${local.pantheon_data_dir}/pantheon/public-keys/",
    "echo Exporting pantheon public-keys",
    "${local.pantheon_binary} public-key export --to=${local.pantheon_data_dir}/pantheon/public-keys/$(hostname)_pubkey",
    "${local.pantheon_binary} public-key export-address --to=${local.pantheon_data_dir}/pantheon/public-keys/$(hostname)_address",

    "export ACCOUNT_ADDRESS=$(cat ${local.pantheon_data_dir}/pantheon/public-keys/$(hostname)_address)",
    "echo Writing account address $ACCOUNT_ADDRESS to ${local.account_address_file}",
    "echo $ACCOUNT_ADDRESS | sed 's/^0x//' > ${local.account_address_file}",

    "export NODE_ID=$(cat ${local.pantheon_data_dir}/pantheon/public-keys/$(hostname)_pubkey)",
    "echo Writing Node Id [$NODE_ID] to ${local.node_id_file}",
    "echo $NODE_ID | sed 's/^0x//' > ${local.node_id_file}",
  ]

  node_key_bootstrap_container_definition = {
    name      = "${local.node_key_bootstrap_container_name}"
    image     = "${local.pantheon_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.pantheon.name}"
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

    cpu = 0
  }

  metadata_bootstrap_commands = [
    "set -e",
    "echo Wait until Node Key is ready ...",
    "while [ ! -f \"${local.node_id_file}\" ]; do sleep 1; done",
    "apk update",
    "apk add curl jq",
    "sleep 10",
    "export TASK_REVISION=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq '.Revision' -r)",
    "echo \"Task Revision: $TASK_REVISION\"",
    "echo $TASK_REVISION > ${local.task_revision_file}",
    "export HOST_IP=$(/usr/bin/curl http://169.254.169.254/latest/meta-data/public-ipv4)",
    "echo \"Host IP: $HOST_IP\"",
    "echo $HOST_IP > ${local.host_ip_file}",
    "export TASK_ARN=$(curl --connect-timeout 5 --retry 5 --max-time 10 -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')",
    "export REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')",
    "aws ecs describe-tasks --region $REGION --cluster ${local.ecs_cluster_name} --tasks $TASK_ARN | jq -r '.tasks[0] | .group' > ${local.service_file}",
    "mkdir -p ${local.hosts_folder}",
    "mkdir -p ${local.node_ids_folder}",
    "mkdir -p ${local.accounts_folder}",
    "mkdir -p ${local.libfaketime_folder}",
    "count=0; while [ $count -lt 1 ]; do count=$(ls ${local.libfaketime_folder} | grep libfaketime.so | wc -l); aws s3 cp s3://${local.s3_libfaketime_file} ${local.libfaketime_folder}/libfaketime.so > /dev/null 2>&1 | echo \"Wait for libfaketime to appear on S3 ... \"; sleep 1; done",
    "touch ${local.libfaketime_file}",
    "aws sqs --region $REGION receive-message --queue-url ${aws_sqs_queue.faketime_queue.id} --visibility-timeout=300 | jq .Messages[].Body | tr -d '\\\"' > ${local.libfaketime_file}",
    "aws s3 cp ${local.node_id_file} s3://${local.s3_revision_folder}/nodeids/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "aws s3 cp ${local.host_ip_file} s3://${local.s3_revision_folder}/hosts/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "aws s3 cp ${local.account_address_file} s3://${local.s3_revision_folder}/accounts/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",

    // Gather all IPs
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.hosts_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/hosts ${local.hosts_folder} > /dev/null 2>&1 | echo \"Wait for other containers to report their IPs ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All containers have reported their IPs\"",

    // Gather all Accounts
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.accounts_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/accounts ${local.accounts_folder} > /dev/null 2>&1 | echo \"Wait for other nodes to report their accounts ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All nodes have registered accounts\"",

    // Gather all Node IDs
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.node_ids_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/nodeids ${local.node_ids_folder} > /dev/null 2>&1 | echo \"Wait for other nodes to report their IDs ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All nodes have registered their IDs\"",

    // Prepare Genesis file
    "alloc=\"\"; for f in $(ls ${local.accounts_folder}); do address=$(cat ${local.accounts_folder}/$f); alloc=\"$alloc,\\\"$address\\\": { \"balance\": \"\\\"1000000000000000000000000000\\\"\"}\"; done",

    "alloc=\"{$${alloc:1}}\"",
    "extraData=\"\\\"RLP_EXTRA_DATA\\\"\"",
    "echo '${replace(jsonencode(local.genesis), "/\"(true|false|[0-9]+)\"/", "$1")}' | jq \". + { alloc : $alloc, extraData: $extraData }\" > ${local.genesis_file}",
    "cat ${local.genesis_file}",

    // Write status
    "echo \"Done!\" > ${local.metadata_bootstrap_container_status_file}",

    "echo Wait until pntheon initialized ...",
    "while [ ! -f \"${local.pantheon_static_nodes_file}\" ]; do sleep 1; done",
    "aws s3 cp ${local.account_address_file} s3://${local.s3_revision_folder}/privacyaddresses/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "echo Done",

  ]

  metadata_bootstrap_container_definition = {
    name      = "${local.metadata_bootstrap_container_name}"
    image     = "${local.aws_cli_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.pantheon.name}"
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

    cpu = 0
  }
}
