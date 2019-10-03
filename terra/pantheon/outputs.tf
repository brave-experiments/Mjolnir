output "_status" {
  value = <<MSG
Completed!

Patnheon Docker Image       = ${local.pantheon_docker_image}
Number of Pantheon Nodes    = ${var.number_of_nodes}
ECS Task Revision           = ${aws_ecs_task_definition.pantheon.revision}
CloudWatch Log Group        = ${aws_cloudwatch_log_group.pantheon.name}
MSG
}

output "bastion_host_ip" {
  value = "${aws_instance.bastion.public_ip}"
}

output "network_name" {
  value = "${var.network_name}"
}

output "ecs_cluster_name" {
  value = "${aws_ecs_cluster.pantheon.name}"
}

output "chain_id" {
  value = "${random_integer.network_id.result}"
}

output "private_key_file" {
  value = "${local_file.private_key.filename}"
  sensitive = true
}

output "ethstats_host_url" {
  value = "http://${aws_instance.bastion.public_ip}:3000"
}

output "bucket_name" {
  value = "${aws_s3_bucket.pantheon.bucket}"
}

output "grafana_host_url" {
  value = "http://${aws_instance.bastion.public_ip}:3001"
}

output "grafana_password" {
  value = "${random_string.random.result}"
}

output "grafana_username" {
  value = "admin"
}
