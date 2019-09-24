resource "aws_security_group" "parity" {
  vpc_id      = "${local.vpc_id}"
  name        = "parity-sg-${var.network_name}"
  description = "Security group used in Parity network ${var.network_name}"

  egress {
    from_port = 0
    protocol  = "-1"
    to_port   = 0

    cidr_blocks = [
      "0.0.0.0/0",
    ]

    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("parity-sg-%s", var.network_name)))}"
}

resource "aws_security_group_rule" "ethstats" {
  from_port         = "${local.ethstats_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.parity.id}"
  to_port           = "${local.ethstats_port}"
  type              = "ingress"
  self              = true
  description       = "ethstats traffic"
}

resource "aws_security_group_rule" "ethstats-bastion" {
  from_port         = "${local.ethstats_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.bastion-ethstats.id}"
  to_port           = "${local.ethstats_port}"
  type              = "ingress"
  self              = true
  description       = "ethstats traffic"
}

resource "aws_security_group_rule" "geth_p2p" {
  from_port         = "${local.parity_p2p_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.parity.id}"
  to_port           = "${local.parity_p2p_port}"
  type              = "ingress"
  self              = true
  description       = "Geth P2P traffic"
}

resource "aws_security_group_rule" "geth_admin_rpc" {
  from_port         = "${local.parity_rpc_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.parity.id}"
  to_port           = "${local.parity_rpc_port}"
  type              = "ingress"
  self              = "true"
  description       = "Geth Admin RPC traffic"
}


resource "aws_security_group_rule" "open-all-ingress-research" {
  count             = "${var.ecs_mode == "EC2" ? 1 : 0}"
  from_port         = 0
  protocol          = "-1"
  security_group_id = "${aws_security_group.parity.id}"
  to_port           = 0
  type              = "ingress"
  cidr_blocks       = ["${var.access_ec2_nodes_cidr_blocks}"]
  description       = "Open all ports"
}
