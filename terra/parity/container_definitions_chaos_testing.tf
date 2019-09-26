locals {

  chaos_testing_run_commands    = "${var.chaos_testing_run_command}"

  chaos_testing_run_container_definition = {
    name      = "${local.chaos_testing_run_container_name}"
    image     = "${local.chaos_testing_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.parity.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
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
    /*
    entrypoint = [
      "sleep",
      "180;",
      "/pumba",
    ]
    */

    command = "${local.chaos_testing_run_commands}"
    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
}
