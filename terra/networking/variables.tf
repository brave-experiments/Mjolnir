# Variables definition

variable "vpc_azs" { type = "list" }
variable vpc_private_subnets { type = "list" }
variable vpc_public_subnets { type = "list" }

variable vpc_cidr {}
#variable vpc_region {}
variable vpc_enable_nat_gateway {}
variable vpc_enable_vpn_gateway {}
variable environment_name {}
variable bucket {}
variable region {}
variable profile {}

