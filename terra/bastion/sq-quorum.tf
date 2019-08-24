resource "aws_security_group" "bastion-ssh" {
  vpc_id      = "${module.vpc.vpc_id}"
  name = "quorum-bastion-ssh-${var.network_name}"
  description = "Security group used by Bastion node to access Quorum network ${var.network_name}"

  ingress {
    from_port = 22
    protocol = "tcp"
    to_port = 22

    cidr_blocks = [
      "${var.access_bastion_cidr_blocks}"
    ]

    description = "Allow SSH"
  }

  egress {
    from_port = 0
    protocol = "-1"
    to_port = 0
    cidr_blocks = [
      "0.0.0.0/0"]
    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("quorum-bastion-ssh-%s", var.network_name)))}"
}

resource "aws_security_group" "bastion-ethstats" {
  vpc_id      = "${module.vpc.vpc_id}"
  name = "quorum-bastion-ethstats-${var.network_name}"
  description = "Security group used by external to access ethstats for Quorum network ${var.network_name}"

  ingress {
    from_port = 3000
    protocol = "tcp"
    to_port = 3000

    cidr_blocks = [
      "${var.access_bastion_cidr_blocks}"
    ]

    description = "Allow ethstats"
  }

  egress {
    from_port = 0
    protocol = "-1"
    to_port = 0
    cidr_blocks = [
      "0.0.0.0/0"]
    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("quorum-bastion-ethstats-%s", var.network_name)))}"
}
