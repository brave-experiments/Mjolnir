locals {
  shared_volume_name             = "parity_shared_volume"
  shared_volume_container_path   = "/qdata"

  node_key_bootstrap_container_name           = "node-key-bootstrap"
  metadata_bootstrap_container_name           = "metamain-bootstrap"
  parity_run_container_name                   = "parity-run"
  chaos_testing_run_container_name            = "chaos-testing-pumba-run"

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
