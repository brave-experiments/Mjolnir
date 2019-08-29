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

  //tags = ["${local.common_tags}"]
  
  user_data = <<EOF
#!/bin/bash

set -e
cd /tmp
curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz
mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/
useradd -rs /bin/false node_exporter


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

#set -e

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
EOF


}
