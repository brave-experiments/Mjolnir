locals {
  pantheon_rpc_port                = 8546
  pantheon_p2p_port                = 30304
  metrics_port                     = 9545
  pantheon_data_dir                = "${local.shared_volume_container_path}/dd"
  pantheon_password_file           = "${local.shared_volume_container_path}/passwords.txt"
  pantheon_static_nodes_file       = "${local.pantheon_data_dir}/static-nodes.json"
  pantheon_permissioned_nodes_file = "${local.pantheon_data_dir}/permissioned-nodes.json"
  genesis_file                   = "${local.shared_volume_container_path}/genesis.json"
  node_id_file                   = "${local.shared_volume_container_path}/node_id"
  node_ids_folder                = "${local.shared_volume_container_path}/nodeids"
  accounts_folder                = "${local.shared_volume_container_path}/accounts"
  privacy_addresses_folder       = "${local.shared_volume_container_path}/privacyaddresses"

  # store Tessera pub keys

  consensus_config_map = "${local.consensus_config[var.consensus_mechanism]}"

  pantheon_config_commands = [
    "mkdir -p ${local.pantheon_data_dir}/geth",
    "echo \"\" > ${local.pantheon_password_file}",
    "echo \"Creating ${local.pantheon_static_nodes_file} and ${local.pantheon_permissioned_nodes_file}\"",
    "all=\"\"; for f in `ls ${local.node_ids_folder}`; do nodeid=$(cat ${local.node_ids_folder}/$f); ip=$(cat ${local.hosts_folder}/$f); all=\"$all,\\\"enode://$nodeid@$ip:${local.pantheon_p2p_port}?discport=0&${join("&", local.consensus_config_map["enode_params"])}\\\"\"; done; all=$${all:1}",
    "echo \"[$all]\" > ${local.pantheon_static_nodes_file}",
    "echo \"[$all]\" > ${local.pantheon_permissioned_nodes_file}",
    "echo Permissioned Nodes: $(cat ${local.pantheon_permissioned_nodes_file})",
    "geth --datadir ${local.pantheon_data_dir} init ${local.genesis_file}",
    "export IDENTITY=$(cat ${local.service_file} | awk -F: '{print $2}')",
  ]

  additional_args = "${local.consensus_config_map["pantheon_args"]}"
  pantheon_args = [
    "--networkid ${random_integer.network_id.result}",
    "--discovery-enabled=false",
    "--p2p-port=30304", 
    "--rpc-http-enabled",
    "--rpc-http-api=ETH,NET,IBFT",
    "--host-whitelist='*'",
    "--rpc-http-cors-origins='all'",
    "--rpc-http-port=8546", 
    "--metrics-enabled",
    "--metrics-host=0.0.0.0",
    "--metrics-port=9545",
    "--host-whitelist=*"]

  pantheon_args_combined = "${join(" ", concat(local.pantheon_args, local.additional_args))}"
  pantheon_run_commands = [
    "set -e",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "echo Wait until ${var.tx_privacy_engine} is ready ...",
    "while [ ! -S \"${local.tx_privacy_engine_socket_file}\" ]; do sleep 1; done",


    "echo 'Running geth with: ${local.pantheon_args_combined}'",
    "/opt/pantheon/bin/pantheon ${local.pantheon_args_combined}",
  ]

  pantheon_run_container_definition = {
    name      = "${local.pantheon_run_container_name}"
    image     = "${local.pantheon_docker_image}"
    essential = "true"

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

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -S ${local.pantheon_data_dir}/geth.ipc ];",
      ]
    }

    environments = []

    //portMappings = []
    portMappings = [
      {
        hostPort      = "${local.pantheon_rpc_port}"
        containerPort = "${local.pantheon_rpc_port}"
      },
      {
        hostPort      = "${local.pantheon_p2p_port}"
        containerPort = "${local.pantheon_p2p_port}"
      },
      {
        hostPort      = "${local.metrics_port}"
        containerPort = "${local.metrics_port}"
      },
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
    ]

    environment = [
      {
        name  = "PRIVATE_CONFIG"
        value = "${local.tx_privacy_engine_socket_file}"
      },
      {
        name  = "LD_PRELOAD",
        value = "${local.libfaketime_folder}/libfaketime.so"
      }
    ]

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.pantheon_run_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
  genesis = {
    "alloc" = {}

    "coinbase" = "0x0000000000000000000000000000000000000000"

    "config" = {
      "homesteadBlock" = 0
      "byzantiumBlock" = 1
      "chainId"        = "${random_integer.network_id.result}"
      "eip150Block"    = 1
      "eip155Block"    = 0
      "eip150Hash"     = "0x0000000000000000000000000000000000000000000000000000000000000000"
      "eip158Block"    = 1
      "ispantheon"       = "true"
    }

    "difficulty" = "${var.genesis_difficulty}"
    "extraData"  = "0x0000000000000000000000000000000000000000000000000000000000000000"
    "gasLimit"   = "${var.genesis_gas_limit}"
    "mixHash"    = "0x00000000000000000000000000000000000000647572616c65787365646c6578"
    "nonce"      = "${var.genesis_nonce}"
    "parentHash" = "0x0000000000000000000000000000000000000000000000000000000000000000"
    "timestamp"  = "${var.genesis_timestamp}"
  }
}

resource "random_integer" "network_id" {
  min = 2018
  max = 9999

  keepers = {
    changes_when = "${var.network_name}"
  }
}
