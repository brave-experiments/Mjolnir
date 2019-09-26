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
  s3_revision_folder = "${local.parity_bucket}/rev_$TASK_REVISION"
  s3_libfaketime_file = "${local.bastion_bucket}/libs/libfaketime.so"
  normalized_host_ip = "ip_$(echo $HOST_IP | sed -e 's/\\./_/g')"

  node_key_bootstrap_commands = [
    "mkdir -p ${local.parity_data_dir}/parity",
    "echo \"\" > ${local.parity_password_file}",
    "bootnode -genkey ${local.parity_data_dir}/parity/nodekey",
    "export NODE_ID=$(bootnode -nodekey ${local.parity_data_dir}/parity/nodekey -writeaddress)",
    "echo Creating an account for this node",
    "geth --datadir ${local.parity_data_dir} account new --password ${local.parity_password_file}",
    "export KEYSTORE_FILE=$(ls ${local.parity_data_dir}/keystore/ | head -n1)",
    "export ACCOUNT_ADDRESS=$(cat ${local.parity_data_dir}/keystore/$KEYSTORE_FILE | sed 's/^.*\"address\":\"\\([^\"]*\\)\".*$/\\1/g')",
    "echo Writing account address $ACCOUNT_ADDRESS to ${local.account_address_file}",
    "echo $ACCOUNT_ADDRESS > ${local.account_address_file}",
    "echo Writing Node Id [$NODE_ID] to ${local.node_id_file}",
    "echo $NODE_ID > ${local.node_id_file}",
  ]

  node_key_bootstrap_container_definition = {
    name      = "${local.node_key_bootstrap_container_name}"
    image     = "${local.quorum_docker_image}"
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

    cpu = 0
  }

  // this is very BADDDDDD but for now i don't have any other better option
  validator_address_program = <<EOP
package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing enode value")
		os.Exit(1)
	}
	enode := os.Args[1]
	nodeId, err := discover.HexID(enode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pub, err := nodeId.Pubkey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("0x%s\n", hex.EncodeToString(crypto.PubkeyToAddress(*pub).Bytes()))
}
EOP

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
    "export TASK_ARN=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')",
    "export REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')",
    "aws ecs describe-tasks --region $REGION --cluster ${local.ecs_cluster_name} --tasks $TASK_ARN | jq -r '.tasks[0] | .group' > ${local.service_file}",
    "mkdir -p ${local.hosts_folder}",
    "mkdir -p ${local.node_ids_folder}",
    "mkdir -p ${local.accounts_folder}",
    "mkdir -p ${local.libfaketime_folder}",
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
    "validators=\"\"; for f in $(ls ${local.accounts_folder}); do address=$(cat ${local.accounts_folder}/$f); validators=\"$validators,\\\"$address\\\" \"; done",
    "accounts=\"\"; for f in $(ls ${local.accounts_folder}); do address=$(cat ${local.accounts_folder}/$f); accounts=\"$accounts,\\\"$address\\\": { \"\\\"balance\\\"\": \"\\\"1000000000000000000000000000\\\"\"}\"; done",

    "validators=\"[ $${validators:1} ]\"",
    "accounts=\"{ $${accounts:1} }\"",
    "echo '${replace(jsonencode(local.genesis), "/\"(true|false|[0-9]+)\"/", "$1")}' | jq \".engine.authorityRound.params.validators.list=$validators\" | jq \".accounts=$accounts\"  > ${local.genesis_file}",
    "sed -i s'/RANDOM_NETWORK_ID/${random_integer.network_id.result}/' ${local.genesis_file}",
    "cat ${local.genesis_file}",

    // Write status
    "echo \"Done!\" > ${local.metadata_bootstrap_container_status_file}",
    "chown 1000:1000 -R ${local.shared_volume_container_path}",

    "echo Wait until parity initialized ...",
    "while [ ! -f \"${local.parity_permissioned_nodes_file}\" ]; do sleep 1; done",
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
