# Variables list

vpc_azs                           = ["us-east-2a", "us-east-2b", "us-east-2c"]
vpc_private_subnets               = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
vpc_public_subnets                = ["10.0.0.0/24"]

vpc_cidr                          = "10.0.0.0/16"
vpc_region                        = "us-east-2"
vpc_enable_nat_gateway            = false
vpc_enable_vpn_gateway            = false
environment_name                  = "sidechain-sandbox"
