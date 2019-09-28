provider "null" {
  version = "~> 1.0"
}

provider "random" {
  version = "~> 2.0"
}

provider "local" {
  version = "~> 1.1"
}

provider "tls" {
  version = "~> 1.2"
}

locals {
  pantheon_docker_image          = "${format("%s:%s", var.pantheon_docker_image, var.pantheon_docker_image_tag)}"
  quorum_docker_image            = "${format("%s:%s", var.quorum_docker_image, var.quorum_docker_image_tag)}"
  aws_cli_docker_image           = "${format("%s:%s", var.aws_cli_docker_image, var.aws_cli_docker_image_tag)}"
  chaos_testing_docker_image     = "${format("%s:%s", var.chaos_testing_docker_image, var.chaos_testing_docker_image_tag)}"

  common_tags = {
    "NetworkName"               = "${var.network_name}"
    "ECSClusterName"            = "${local.ecs_cluster_name}"
    "DockerImage.Pantheon"      = "${local.pantheon_docker_image}"
  }
}
