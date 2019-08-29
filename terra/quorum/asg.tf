module "ecs-asg" {
  source  = "silinternational/ecs-asg/aws"
  version = "1.1.1"

  #required
  cluster_name    = "${local.ecs_cluster_name}"
  security_groups = ["${aws_security_group.quorum.id}"]
  subnet_ids      = ["${module.vpc.public_subnets[0]}"]

  #optional
  alarm_actions_enabled = false
  adjustment_type       = "ExactCapacity"
  instance_type         = "${var.asg_instance_type}"
  max_size              = "${var.number_of_nodes}"
  min_size              = "${var.number_of_nodes}"

  # blocks --destroy option
  protect_from_scale_in = false
  root_volume_size      = "16"
  ssh_key_name          = "${aws_key_pair.ssh.key_name}"

  user_data       = "${file("${path.module}/files/node_exporter_install")}"

  //tags = ["${local.common_tags}"]
}
