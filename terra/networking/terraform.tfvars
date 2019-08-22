# Variables list

vpc_azs                           = ["us-east-2a", "us-east-2b", "us-east-2c"]
vpc_private_subnets               = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
vpc_public_subnets                = ["10.0.0.0/24"]

vpc_cidr                          = "10.0.0.0/16"
vpc_enable_nat_gateway            = false
vpc_enable_vpn_gateway            = false
environment_name                  = "${var.network_name}"

region                            = "us-east-2"
#TODO cover creating bucket and KMS by tf
bucket                            = "us-east-2-quorum-deployment-state-051582052996"
encrypt                           = "true"
kms_key_id                        = "arn:aws:kms:us-east-2:051582052996:key/6b957713-73ed-4795-9fc5-a77844af8407"

#TODO pass it from CLI
profile                           = "binarapps-brave-sidechain-sandbox"
