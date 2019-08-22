terraform {
  backend "s3" {
    # backend configuration is auto discovered by running Terraform inside _terraform_init folder
  }
}

provider "aws" {
  region  = "${var.region}"
  version = "~> 1.36"
  profile = "${var.profile}"
}

data "aws_security_group" "default" {
  name   = "default"
  vpc_id = "${module.vpc.vpc_id}"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "~> 1.60"

  name = "${var.environment_name}"
  cidr = "${var.vpc_cidr}"

  azs             = "${var.vpc_azs}"
  private_subnets = "${var.vpc_private_subnets}"
  public_subnets  = "${var.vpc_public_subnets}"

  enable_nat_gateway = "${var.vpc_enable_nat_gateway}"
  enable_vpn_gateway = "${var.vpc_enable_vpn_gateway}"

  tags = {
    Terraform = "true"
    Environment = "${var.environment_name}"
  }
}
