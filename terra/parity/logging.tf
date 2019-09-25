resource "aws_cloudwatch_log_group" "parity" {
  name              = "/ecs/parity/${var.network_name}"
  retention_in_days = "7"
  tags              = "${local.common_tags}"
}
