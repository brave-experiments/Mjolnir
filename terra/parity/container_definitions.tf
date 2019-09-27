locals {
  shared_volume_name             = "parity_shared_volume"
  shared_volume_container_path   = "/qdata"
  tx_privacy_engine_socket_file  = "${local.shared_volume_container_path}/tm.ipc"
  tx_privacy_engine_address_file = "${local.shared_volume_container_path}/tm.pub"

  node_key_bootstrap_container_name           = "node-key-bootstrap"
  metadata_bootstrap_container_name           = "metamain-bootstrap"
  parity_run_container_name                   = "parity-run"
  chaos_testing_run_container_name            = "chaos-testing-pumba-run"

  consensus_config = {
    poa = {
      parity_args = [
      ]

      enode_params = []

      genesis_mixHash    = ["0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365"]
      genesis_difficulty = ["0x01"]

      git_url = ["https://github.com/getamis/istanbul-tools"]
    }
  }

  common_container_definitions = [
    "${local.node_key_bootstrap_container_definition}",
    "${local.metadata_bootstrap_container_definition}",
    "${local.parity_run_container_definition}",
    "${local.chaos_testing_run_container_definition}",
  ]

  container_definitions = [
    "${jsonencode(local.common_container_definitions)}",

  ]
}
