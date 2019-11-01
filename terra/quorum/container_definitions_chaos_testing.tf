locals {

  chaos_testing_run_commands    = "${formatlist("%s", var.chaos_testing_run_command)}"

  chaos_testing_run_container_definition = {
    name      = "${local.chaos_testing_run_container_name}"
    image     = "${local.chaos_testing_docker_image}"
    essential = "false"
    memoryReservation = "8192"

    logConfiguration = {
      logDriver = "fluentd"

      options = {
        fluentd-address  = "${aws_instance.bastion.public_ip}:24224"
      }
    }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
      {
        sourceVolume  = "docker_socket"
        containerPath = "/var/run/docker.sock"
      },
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
    ]

    entrypoint = "${concat(list("/pumba"), local.chaos_testing_run_commands)}"
    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
}
