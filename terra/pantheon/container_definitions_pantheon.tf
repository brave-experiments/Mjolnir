locals {
  pantheon_rpc_port                = 8545
  pantheon_ws_port                 = 8546
  pantheon_p2p_port                = 30303
  pantheon_metrics_port            = 9545
  pantheon_data_dir                = "${local.shared_volume_container_path}/dd"
  pantheon_binary                  = "/opt/pantheon/bin/pantheon --data-path=${local.pantheon_data_dir}"
  pantheon_password_file           = "${local.shared_volume_container_path}/passwords.txt"
  pantheon_static_nodes_file       = "${local.pantheon_data_dir}/static-nodes.json"
  pantheon_permissioned_nodes_file = "${local.pantheon_data_dir}/permissioned-nodes.json"
  genesis_file                   = "${local.shared_volume_container_path}/genesis.json"
  node_id_file                   = "${local.shared_volume_container_path}/node_id"
  node_ids_folder                = "${local.shared_volume_container_path}/nodeids"
  accounts_folder                = "${local.shared_volume_container_path}/accounts"
  privacy_addresses_folder       = "${local.shared_volume_container_path}/privacyaddresses"
  faketime_dont_fake_monotonic   = 1


  pantheon_config_commands = [
    "echo \"\" > ${local.pantheon_password_file}",
    "echo \"Creating ${local.pantheon_static_nodes_file} and ${local.pantheon_permissioned_nodes_file}\"",

    "all=\"\"; for f in $(ls ${local.node_ids_folder}); do nodeid=$(cat ${local.node_ids_folder}/$f); ip=$(cat ${local.hosts_folder}/$f); all=\"$all,\\\"enode://$nodeid@$ip:${local.pantheon_p2p_port}\\\"\"; done; ",
    "echo \"[ $(echo $all | sed 's/^.//') ] \" > ${local.pantheon_static_nodes_file}",
    "unset all",

    "echo \"Creating Encode.json Validators list\"",
    "export FAKETIME_DONT_FAKE_MONOTONIC=1",
    "all=\"\"; for f in $(ls ${local.node_ids_folder}); do address=$(cat ${local.accounts_folder}/$f); all=\"$all,\\\"$address\\\"\"; done;  ",  
    "echo \"[ $(echo $all | sed 's/^.//') ] \" > toEncode.json",
    "cat ${local.pantheon_static_nodes_file}",
    "cat toEncode.json",

    # replace placeholder by encoded rpl address list in genesis
    "export rlp=$(${local.pantheon_binary} rlp encode --from=toEncode.json)",
    "echo $rlp",
    "sed -i s/RLP_EXTRA_DATA/$rlp/ ${local.genesis_file}",
    "cat ${local.genesis_file}",
    "cp ${local.pantheon_static_nodes_file} ${local.pantheon_permissioned_nodes_file}",
    "cp ${local.pantheon_static_nodes_file} static-nodes.json",
    "echo Permissioned Nodes: $(cat ${local.pantheon_permissioned_nodes_file})",

  ]

  pantheon_args = [
    "--genesis-file=${local.genesis_file}",
    "--network-id=${random_integer.network_id.result}",
    "--discovery-enabled=false",
    "--p2p-port=${local.pantheon_p2p_port}", 
    "--rpc-http-enabled",
    "--rpc-http-api=WEB3,ETH,NET,IBFT,ADMIN",
    "--rpc-http-host=0.0.0.0",
    "--rpc-http-port=${local.pantheon_rpc_port}", 
    "--rpc-http-cors-origins=*",
    "--rpc-ws-enabled",
    "--rpc-ws-api=WEB3,ETH,NET,IBFT",
    "--rpc-ws-host=0.0.0.0",
    "--rpc-ws-port=${local.pantheon_ws_port}",
    "--metrics-enabled",
    "--metrics-host=0.0.0.0",
    "--metrics-port=${local.pantheon_metrics_port}",
    "--host-whitelist=*"
    ]

  pantheon_args_combined = "${join(" ", local.pantheon_args)}"
  pantheon_run_commands = [
    "set -e",
    "count=0; while [ $count -lt 1 ]; do count=$(ls ${local.libfaketime_folder} | grep libfaketime.so | wc -l); echo \"Wait for libfaketime to appear on storage ... \"; sleep 10; done",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "echo Metadata bootstrap completed",

    "${local.pantheon_config_commands}",

    "echo 'Running pantheon with: ${local.pantheon_args_combined}'",
    "${local.pantheon_binary} ${local.pantheon_args_combined}",
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
        "[ -f ${local.genesis_file} ];",
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
        hostPort      = "${local.pantheon_ws_port}"
        containerPort = "${local.pantheon_ws_port}"
      },
      {
        hostPort      = "${local.pantheon_p2p_port}"
        containerPort = "${local.pantheon_p2p_port}"
      },
      {
        hostPort      = "${local.pantheon_metrics_port}"
        containerPort = "${local.pantheon_metrics_port}"
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
        name  = "LD_PRELOAD"
        value = "${local.libfaketime_folder}/libfaketime.so"
      },
      {
        name  = "FAKETIME_TIMESTAMP_FILE"
        value = "${local.libfaketime_file}"
      },
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
      "chainId"                   = "${random_integer.network_id.result}"
      "constantinoplefixblock"    = 0
      "ibft2" = {
        "blockperiodseconds"      = "${var.genesis_block_period_seconds}"
        "epochlength"             = 30000
        "requesttimeoutseconds"   = 10
      }
    },

    "difficulty" = "${var.genesis_difficulty}"
    "extraData"  = "RLP_EXTRA_DATA"
    "gasLimit"   = "${var.genesis_gas_limit}"
    "mixHash"    = "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365"
    "nonce"      = "${var.genesis_nonce}"
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
