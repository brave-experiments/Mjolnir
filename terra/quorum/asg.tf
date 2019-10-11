/*
 * Generate user_data from template file
 */
data "template_file" "user_data" {
  template = <<EOT
    #!/bin/bash
    echo ECS_CLUSTER=${local.ecs_cluster_name} >> /etc/ecs/ecs.config

    # node_exporter part
    set -e
    cd /tmp
    curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
    tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz
    mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/
    useradd -rs /bin/false node_exporter
    echo "ECS_AVAILABLE_LOGGING_DRIVERS=[\"awslogs\",\"fluentd\"]" >> /etc/ecs/ecs.config


    tee -a /etc/init.d/node_exporter << END
#!/bin/bash

### BEGIN INIT INFO
# processname:       node_exporter
# Short-Description: Exporter for machine metrics.
# Description:       Prometheus exporter for machine metrics,
#                    written in Go with pluggable metric collectors.
#
# chkconfig: 2345 80 80
# pidfile: /var/run/node_exporter/node_exporter.pid
#
#
### END INIT INFO

# Source function library.
. /etc/init.d/functions

NAME=node_exporter
DESC="Exporter for machine metrics"
DAEMON=/usr/local/bin/node_exporter
USER=node_exporter
CONFIG=
PID=/var/run/node_exporter/\$NAME.pid
LOG=/var/log/node_exporter/\$NAME.log

DAEMON_OPTS=
RETVAL=0

# Check if DAEMON binary exist
[ -f \$DAEMON ] || exit 0

[ -f /etc/default/node_exporter ]  &&  . /etc/default/node_exporter

service_checks() {
  # Prepare directories
  mkdir -p /var/run/node_exporter /var/log/node_exporter
  chown -R \$USER /var/run/node_exporter /var/log/node_exporter

  # Check if PID exists
  if [ -f "\$PID" ]; then
    PID_NUMBER=\$(cat \$PID)
    if [ -z "\$(ps axf | grep \$PID_NUMBER | grep -v grep)" ]; then
      echo "Service was aborted abnormally; clean the PID file and continue..."
      rm -f "\$PID"
    else
      echo "Service already started; skip..."
      exit 1
    fi
  fi
}

start() {
  service_checks \$1
  sudo -H -u \$USER   \$DAEMON \$DAEMON_OPTS  > \$LOG 2>&1  &
  RETVAL=\$?
  echo \$! > \$PID
}

stop() {
  killproc -p \$PID -b \$DAEMON  \$NAME
  RETVAL=\$?
}

reload() {
  #-- sorry but node_exporter doesn't handle -HUP signal...
  #killproc -p \$PID -b \$DAEMON  \$NAME -HUP
  #RETVAL=\$?
  stop
  start
}

case "\$1" in
  start)
    echo -n \$"Starting \$DESC -" "\$NAME" \$'\n'
    start
    ;;

  stop)
    echo -n \$"Stopping \$DESC -" "\$NAME" \$'\n'
    stop
    ;;

  reload)
    echo -n \$"Reloading \$DESC configuration -" "\$NAME" \$'\n'
    reload
    ;;

  restart|force-reload)
    echo -n \$"Restarting \$DESC -" "\$NAME" \$'\n'
    stop
    start
    ;;

  syntax)
    \$DAEMON --help
    ;;

  status)
    status -p \$PID \$DAEMON
    ;;

  *)
    echo -n \$"Usage: /etc/init.d/\$NAME {start|stop|reload|restart|force-reload|syntax|status}" \$'\n'
    ;;
esac

exit \$RETVAL
END

chmod +x /etc/init.d/node_exporter
service node_exporter start
chkconfig node_exporter on

EOT

  vars {
    ecs_cluster_name = "${local.ecs_cluster_name}"
  }
}

/*
 * Create Launch Configuration
 */
resource "aws_launch_configuration" "lc" {
  image_id             = "${data.aws_ami.ecs_ami.id}"
  name_prefix          = "${local.ecs_cluster_name}"
  instance_type        = "${var.asg_instance_type}"
  iam_instance_profile = "${aws_iam_instance_profile.ecsInstanceProfile.id}"
  security_groups      = ["${aws_security_group.quorum.id}"]
  user_data            = "${var.user_data != "false" ? var.user_data : data.template_file.user_data.rendered}"
  key_name             = "${aws_key_pair.ssh.key_name}"

  root_block_device {
    volume_size = "${var.root_volume_size}"
  }

  lifecycle {
    create_before_destroy = true
  }
}

/*
 * Create Auto-Scaling Group
 */
resource "aws_autoscaling_group" "asg" {
  name                      = "${local.ecs_cluster_name}"
  vpc_zone_identifier       = ["${aws_subnet.public.*.id}"]
  min_size                  = "${var.number_of_nodes}"
  max_size                  = "${var.number_of_nodes}"
  health_check_type         = "${var.health_check_type}"
  health_check_grace_period = "${var.health_check_grace_period}"
  default_cooldown          = "${var.default_cooldown}"
  termination_policies      = ["${var.termination_policies}"]
  launch_configuration      = "${aws_launch_configuration.lc.id}"

  tags = ["${concat(
    list(
      map("key", "ecs_cluster", "value", local.ecs_cluster_name, "propagate_at_launch", true)
    ),
    var.asg_tags
  )}"]

  protect_from_scale_in = "${var.protect_from_scale_in}"

  lifecycle {
    create_before_destroy = true
  }
}

/*
 * Create autoscaling policies
 */
resource "aws_autoscaling_policy" "up" {
  name                   = "${local.ecs_cluster_name}-scaleUp"
  scaling_adjustment     = "${var.scaling_adjustment_up}"
  adjustment_type        = "${var.adjustment_type}"
  cooldown               = "${var.policy_cooldown}"
  policy_type            = "SimpleScaling"
  autoscaling_group_name = "${aws_autoscaling_group.asg.name}"
  count                  = "${var.alarm_actions_enabled ? 1 : 0}"
}

resource "aws_autoscaling_policy" "down" {
  name                   = "${local.ecs_cluster_name}-scaleDown"
  scaling_adjustment     = "${var.scaling_adjustment_down}"
  adjustment_type        = "${var.adjustment_type}"
  cooldown               = "${var.policy_cooldown}"
  policy_type            = "SimpleScaling"
  autoscaling_group_name = "${aws_autoscaling_group.asg.name}"
  count                  = "${var.alarm_actions_enabled ? 1 : 0}"
}

/*
 * Create CloudWatch alarms to trigger scaling of ASG
 */
resource "aws_cloudwatch_metric_alarm" "scaleUp" {
  alarm_name          = "${local.ecs_cluster_name}-scaleUp"
  alarm_description   = "ECS cluster scaling metric above threshold"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = "${var.evaluation_periods}"
  metric_name         = "${var.scaling_metric_name}"
  namespace           = "AWS/ECS"
  statistic           = "Average"
  period              = "${var.alarm_period}"
  threshold           = "${var.alarm_threshold_up}"
  actions_enabled     = "${var.alarm_actions_enabled}"
  count               = "${var.alarm_actions_enabled ? 1 : 0}"
  alarm_actions       = ["${aws_autoscaling_policy.up.arn}"]

  dimensions {
    ClusterName = "${local.ecs_cluster_name}"
  }
}

resource "aws_cloudwatch_metric_alarm" "scaleDown" {
  alarm_name          = "${local.ecs_cluster_name}-scaleDown"
  alarm_description   = "ECS cluster scaling metric under threshold"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "${var.evaluation_periods}"
  metric_name         = "${var.scaling_metric_name}"
  namespace           = "AWS/ECS"
  statistic           = "Average"
  period              = "${var.alarm_period}"
  threshold           = "${var.alarm_threshold_down}"
  actions_enabled     = "${var.alarm_actions_enabled}"
  count               = "${var.alarm_actions_enabled ? 1 : 0}"
  alarm_actions       = ["${aws_autoscaling_policy.down.arn}"]

  dimensions {
    ClusterName = "${local.ecs_cluster_name}"
  }
}

variable "user_data" {
  description = "Bash code for inclusion as user_data on instances. By default contains minimum for registering with ECS cluster"
  default     = "false"
}

variable "root_volume_size" {
  default = "16"
}

variable "min_size" {
  default = "1"
}

variable "max_size" {
  default = "5"
}

variable "health_check_type" {
  default = "EC2"
}

variable "health_check_grace_period" {
  default = "300"
}

variable "default_cooldown" {
  default = "30"
}

variable "termination_policies" {
  type        = "list"
  default     = ["Default"]
  description = "The allowed values are OldestInstance, NewestInstance, OldestLaunchConfiguration, ClosestToNextInstanceHour, Default."
}

variable "protect_from_scale_in" {
  default = false
}

variable "asg_tags" {
  type        = "list"
  description = "List of maps with keys: 'key', 'value', and 'propagate_at_launch'"

  default = [
    {
      key                 = "created_by"
      value               = "terraform"
      propagate_at_launch = true
    },
  ]
}

variable "scaling_adjustment_up" {
  default     = "1"
  description = "How many instances to scale up by when triggered"
}

variable "scaling_adjustment_down" {
  default     = "-1"
  description = "How many instances to scale down by when triggered"
}

variable "scaling_metric_name" {
  default     = "CPUReservation"
  description = "Options: CPUReservation or MemoryReservation"
}

variable "adjustment_type" {
  default     = "ExactCapacity"
  description = "Options: ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity"
}

variable "policy_cooldown" {
  default     = 300
  description = "The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start."
}

variable "evaluation_periods" {
  default     = "2"
  description = "The number of periods over which data is compared to the specified threshold."
}

variable "alarm_period" {
  default     = "120"
  description = "The period in seconds over which the specified statistic is applied."
}

variable "alarm_threshold_up" {
  default     = "100"
  description = "The value against which the specified statistic is compared."
}

variable "alarm_threshold_down" {
  default     = "50"
  description = "The value against which the specified statistic is compared."
}

variable "alarm_actions_enabled" {
  default = false
}
