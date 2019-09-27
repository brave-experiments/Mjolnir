locals {
  shared_volume_name             = "pantheon_shared_volume"
  shared_volume_container_path   = "/qdata"
 
  node_key_bootstrap_container_name           = "node-key-bootstrap"
  metadata_bootstrap_container_name           = "metamain-bootstrap"
  pantheon_run_container_name                 = "pantheon-run"
  istanbul_extradata_bootstrap_container_name = "ibft-extramain-bootstrap"

  common_container_definitions = [
    "${local.node_key_bootstrap_container_definition}",
    "${local.metadata_bootstrap_container_definition}",
    "${local.pantheon_run_container_definition}",
  ]

  container_definitions = [
    "${jsonencode(local.common_container_definitions)}",
  ]
}
