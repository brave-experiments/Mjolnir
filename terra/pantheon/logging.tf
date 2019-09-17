resource "aws_cloudwatch_log_group" "pantheon" {
  name              = "/ecs/pantheon/${var.network_name}"
  retention_in_days = "7"
  tags              = "${local.common_tags}"
}
