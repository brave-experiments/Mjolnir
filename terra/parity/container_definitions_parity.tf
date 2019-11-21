locals {
  parity_rpc_port                = 8545
  parity_p2p_port                = 30303
  parity_data_dir                = "${local.shared_volume_container_path}/dd"
  parity_password_file           = "${local.shared_volume_container_path}/passwords.txt"
  parity_static_nodes_file       = "${local.parity_data_dir}/static-nodes.json"
  parity_permissioned_nodes_file = "${local.parity_data_dir}/permissioned-nodes.json"
  config_file                    = "${local.config_folder}/node.toml"
  genesis_file                   = "${local.shared_volume_container_path}/spec.json"
  node_id_file                   = "${local.shared_volume_container_path}/node_id"
  keys_folder                    = "${local.shared_volume_container_path}/keys"
  keys_json_folder               = "${local.shared_volume_container_path}/keys_json"
  config_folder                  = "${local.shared_volume_container_path}/config"
  privacy_addresses_folder       = "${local.shared_volume_container_path}/privacyaddresses"
  chain_name                     = "MjolnirPOA"

  parity_config_commands = [
    "mkdir -p ${local.parity_data_dir}/parity",
    "echo \"\" > ${local.parity_password_file}",
    "export FAKETIME_DONT_FAKE_MONOTONIC=1",
    "ls /node",
    "ls /node/config",
    "cp /node/spec.json /node/data/spec.json",
  ]
  parity_args = [
    "--config ${local.config_file}",
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
      "hbbft" = {
        "params" = {
          "minimumBlockTime" = "${var.genesis_blocktime}",
          "transactionQueueSizeTrigger" = 1
        }
      }
    },
    "params" = {
      "gasLimitBoundDivisor" = "0x400",
      "maximumExtraDataSize" = "0x20",
      "minGasLimit" = "${var.genesis_min_gas_limit}",
      "networkID" = "RANDOM_NETWORK_ID",
      "eip140Transition" = "0x0",
      "eip211Transition" = "0x0",
      "eip214Transition" = "0x0",
      "eip658Transition" = "0x0",
      "eip145Transition" = "0x0",
      "eip1014Transition" = "0x0",
      "eip1052Transition" = "0x0"
    },
    "genesis" = {
      "seal" = {
        "generic" = "0x0"
      },
      "difficulty" = "${var.genesis_difficulty}",
      "gasLimit"   = "${var.genesis_gas_limit}",
      "timestamp"  = "${var.genesis_timestamp}"
      "author" = "0x0000000000000000000000000000000000000000",
      "parentHash" = "0x0000000000000000000000000000000000000000000000000000000000000000",
      "extraData" = "0x",
    },
    "accounts" = {
      "0000000000000000000000000000000000000005" = {
        "builtin" = {
          "name" = "modexp",
          "activate_at" = "0x0",
          "pricing" = {
            "modexp" = {
              "divisor" = 20
            }
          }
        }
      },
      "0000000000000000000000000000000000000006" =   {
        "builtin" = {
          "name" = "alt_bn128_add",
          "activate_at" = "0x0",
          "eip1108_transition" = "0x0",
          "pricing" = {
            "alt_bn128_const_operations" = {
              "price" = 500,
              "eip1108_transition_price" = 150
            }
          }
        }
      },
      "0000000000000000000000000000000000000007" = {
        "builtin" = {
          "name" = "alt_bn128_mul",
          "activate_at" = "0x0",
          "eip1108_transition" = "0x0",
          "pricing" = {
            "alt_bn128_const_operations" = {
              "price" = 40000,
              "eip1108_transition_price" = 6000
            }
          }
        }
      },
      "0000000000000000000000000000000000000008" = {
        "builtin" = {
          "name" = "alt_bn128_pairing",
          "activate_at" = "0x0",
          "eip1108_transition" = "0x0",
          "pricing" = {
            "alt_bn128_pairing" = {
              "base" = 100000,
              "pair" = 80000,
              "eip1108_transition_base" = 45000,
              "eip1108_transition_pair" = 34000
            }
          }
        }
      },
      "0x0000000000000000000000000000000000000009" = {
        "builtin" = {
          "name" = "blake2_f",
          "activate_at" = "0x0",
          "pricing" = {
            "blake2_f" = {
              "gas_per_round" = 1
            }
          }
        }
      },
      "0x0000000000000000000000000000000000000001" = {
        "balance" = "1",
        "builtin" = {
          "name" = "ecrecover",
          "pricing" = {
            "linear" = {
              "base" = 3000,
              "word" = 0
            }
          }
        }
      },
      "0x0000000000000000000000000000000000000002" = {
        "balance" = "1",
        "builtin" = {
          "name" = "sha256",
          "pricing" = {
            "linear" = {
              "base" = 60,
              "word" = 12
            }
          }
        }
      },
      "0x0000000000000000000000000000000000000003" = {
        "balance" = "1",
        "builtin" = {
          "name" = "ripemd160",
          "pricing" = {
            "linear" = {
              "base" = 600,
              "word" = 120
            }
          }
        }
      },
      "0x0000000000000000000000000000000000000004" = {
        "balance" = "1",
        "builtin" = {
          "name" = "identity",
          "pricing" = {
            "linear" = {
              "base" = 15,
              "word" = 3
            }
          }
        }
      },
     
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
