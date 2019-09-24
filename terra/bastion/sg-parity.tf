resource "aws_security_group" "bastion-ssh" {
  vpc_id      = "${local.vpc_id}"
  name        = "parity-bastion-ssh-${var.network_name}"
  description = "Security group used by Bastion node to access Parity network ${var.network_name}"

  ingress {
    from_port = 22
    protocol  = "tcp"
    to_port   = 22

    cidr_blocks = [
      "${var.access_bastion_cidr_blocks}",
    ]

    description = "Allow SSH"
  }

  egress {
    from_port = 0
    protocol  = "-1"
    to_port   = 0

    cidr_blocks = [
      "0.0.0.0/0",
    ]

    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("parity-bastion-ssh-%s", var.network_name)))}"
}

resource "aws_security_group" "bastion-ethstats" {
  vpc_id      = "${local.vpc_id}"
  name        = "parity-bastion-ethstats-${var.network_name}"
  description = "Security group used by external to access ethstats for parity network ${var.network_name}"

  ingress {
    from_port = "${local.parity_rpc_port}"
    protocol  = "tcp"
    to_port   = "${local.parity_rpc_port}"

    cidr_blocks = [
      "${aws_subnet.public.cidr_block}",
    ]

    description = "Allow geth console"
  }

  egress {
    from_port = 0
    protocol  = "-1"
    to_port   = 0

    cidr_blocks = [
      "0.0.0.0/0",
    ]

    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("client-bastion-geth-%s", var.network_name)))}"
}

resource "aws_security_group" "bastion-geth" {
  vpc_id      = "${local.vpc_id}"
  name        = "${var.client_name}-bastion-get-${var.network_name}"
  description = "Security group used by external to access geth for ${var.client_name} network ${var.network_name}"

  ingress {
    from_port = 3000
    protocol  = "tcp"
    to_port   = 3000

    cidr_blocks = [
      "${var.access_bastion_cidr_blocks}",
    ]

    description = "Allow ethstats"
  }

  egress {
    from_port = 0
    protocol  = "-1"
    to_port   = 0

    cidr_blocks = [
      "0.0.0.0/0",
    ]

    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("%s-bastion-ethstats-%s", var.client_name ,var.network_name)))}"
}