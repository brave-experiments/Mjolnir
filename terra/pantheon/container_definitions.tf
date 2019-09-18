locals {
  shared_volume_name             = "pantheon_shared_volume"
  shared_volume_container_path   = "/qdata"
  tx_privacy_engine_socket_file  = "${local.shared_volume_container_path}/tm.ipc"
  tx_privacy_engine_address_file = "${local.shared_volume_container_path}/tm.pub"
 
  node_key_bootstrap_container_name           = "node-key-bootstrap"
  metadata_bootstrap_container_name           = "metamain-bootstrap"
  pantheon_run_container_name                 = "pantheon-run"
  istanbul_extradata_bootstrap_container_name = "ibft-extramain-bootstrap"


  consensus_config = {

    istanbul = {
      pantheon_args = []

      enode_params = []

      genesis_mixHash    = ["0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365"]
      genesis_difficulty = ["0x01"]

      git_url = [""]
    }
  }

  common_container_definitions = [
    "${local.node_key_bootstrap_container_definition}",
    "${local.metadata_bootstrap_container_definition}",
    "${local.pantheon_run_container_definition}",
  ]

  container_definitions = [
    "${jsonencode(local.common_container_definitions)}",
  ]
}
