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
  //tx_privacy_engine_docker_image = "${coalesce(local.tessera_docker_image, local.constellation_docker_image)}"
  aws_cli_docker_image           = "${format("%s:%s", var.aws_cli_docker_image, var.aws_cli_docker_image_tag)}"

  common_tags = {
    "NetworkName"               = "${var.network_name}"
    "ECSClusterName"            = "${local.ecs_cluster_name}"
    "DockerImage.Pantheon"      = "${local.pantheon_docker_image}"
  }
}
