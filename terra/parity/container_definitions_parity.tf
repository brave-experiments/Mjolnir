locals {
  parity_rpc_port                = 8545
  parity_p2p_port                = 30303
  parity_data_dir                = "${local.shared_volume_container_path}/dd"
  parity_password_file           = "${local.shared_volume_container_path}/passwords.txt"
  parity_static_nodes_file       = "${local.parity_data_dir}/static-nodes.json"
  parity_permissioned_nodes_file = "${local.parity_data_dir}/permissioned-nodes.json"
  genesis_file                   = "${local.shared_volume_container_path}/genesis.json"
  node_id_file                   = "${local.shared_volume_container_path}/node_id"
  node_ids_folder                = "${local.shared_volume_container_path}/nodeids"
  accounts_folder                = "${local.shared_volume_container_path}/accounts"
  privacy_addresses_folder       = "${local.shared_volume_container_path}/privacyaddresses"
  chain_name                     = "ApolloPOA"

  parity_config_commands = [
    "mkdir -p ${local.parity_data_dir}/parity",
    "echo \"\" > ${local.parity_password_file}",
    "export FAKETIME_DONT_FAKE_MONOTONIC=1",
    "echo \"Creating ${local.parity_static_nodes_file} and ${local.parity_permissioned_nodes_file}\"",
    "for f in $(ls ${local.node_ids_folder}); do nodeid=$(cat ${local.node_ids_folder}/$f); ip=$(cat ${local.hosts_folder}/$f); echo \"enode://$nodeid@$ip:${local.parity_p2p_port}\" >> ${local.parity_static_nodes_file}; done;",
    "cat ${local.parity_static_nodes_file} > ${local.parity_permissioned_nodes_file}",
    "echo Permissioned Nodes: $(cat ${local.parity_permissioned_nodes_file})",
    "export IDENTITY=$(cat ${local.service_file} | awk -F: '{print $2}')",
  ]
  parity_args = [
    "--chain ${local.genesis_file}",
    "--base-path ${local.parity_data_dir}",
    "--keys-path ${local.parity_data_dir}/keystore",
    "--interface 0.0.0.0",
    "--jsonrpc-interface 0.0.0.0",
    "--jsonrpc-apis eth,net,shh,personal,web3,parity,parity_set,parity_accounts,traces,rpc",
    "--jsonrpc-port ${local.parity_rpc_port}",
    "--port ${local.parity_p2p_port}",
    "--password ${local.parity_password_file}",
    "--no-discovery",
    "--reserved-only",
    "--reserved-peers ${local.parity_static_nodes_file}",
    "--engine-signer 0x$(cat ${local.account_address_file})",
    "--force-sealing",
    "--unsafe-expose",
  ]
  parity_args_combined = "${join(" ", local.parity_args)}"
  parity_run_commands = [
    "set -e",
    //"cat ${local.libfaketime_file} > /etc/faketimerc",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "${local.parity_config_commands}",
    "echo 'Running parity with: ${local.parity_args_combined}'",
    "parity ${local.parity_args_combined}",
  ]
  parity_run_container_definition = {
    name      = "${local.parity_run_container_name}"
    image     = "${local.parity_docker_image}"
    essential = "true"

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

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -S ${local.parity_data_dir}/jsonrpc.ipc ];",
      ]
    }

    environments = []

    //portMappings = []
    portMappings = [
      {
        hostPort      = "${local.parity_rpc_port}"
        containerPort = "${local.parity_rpc_port}"
      },
      {
        hostPort      = "${local.parity_p2p_port}"
        containerPort = "${local.parity_p2p_port}"
      }
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
    ]

    environment = [
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
      "${join("\n", local.parity_run_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }

  genesis = {
    "name" = "${local.chain_name}",
    "engine" = {
      "authorityRound" = {
        "params" = {
          "stepDuration" = "${var.genesis_step_duration}",
          "validators" = {
            "list" = [
            ]
          }
        }
      }
    },
    "params" = {
      "gasLimitBoundDivisor" = "0x400",
      "maximumExtraDataSize" = "0x20",
      "minGasLimit" = "${var.genesis_min_gas_limit}",
      "networkID" = "RANDOM_NETWORK_ID",
      "eip155Transition" = 0,
      "validateChainIdTransition" = 0,
      "eip140Transition" = 0,
      "eip211Transition" = 0,
      "eip214Transition" = 0,
      "eip658Transition" = 0
    },
    "genesis" = {
      "seal" = {
        "authorityRound" = {
          "step" = "0x0",
          "signature" = "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
        }
      },
      "difficulty" = "${var.genesis_difficulty}",
      "gasLimit"   = "${var.genesis_gas_limit}",
      "timestamp"  = "${var.genesis_timestamp}"
    },
    "accounts" = {
    }
  }
}

resource "random_integer" "network_id" {
  min = 2018
  max = 9999

  keepers = {
    changes_when = "${var.network_name}"
  }
}
