# Variables list

azs                           = ["us-east-2a", "us-east-2b", "us-east-2c"]
private_subnets               = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
public_subnets                = ["10.0.0.0/24"]

cidr                          = "10.0.0.0/16"
enable_nat_gateway            = false
enable_vpn_gateway            = false

#TODO pass it from CLI
#profile                           = "binarapps-brave-sidechain-sandbox"

is_igw_subnets = "false"

# private subnets routable to Internet via NAT Gateway
subnet_ids = [
    "${module.vpc.public_subnets[0]}",
]

vpc_id = "${module.vpc.vpc_id}"

bastion_public_subnet_id = "${module.vpc.public_subnets[0]}"

consensus_mechanism = "istanbul"

# tx_privacy_engine = "constellation"

access_bastion_cidr_blocks = [
  "0.0.0.0/0",
]

# Open access to ALL PORTS on EC2 cluster nodes
access_ec2_nodes_cidr_blocks = [
  "0.0.0.0/0",
]

number_of_nodes = 0 

# EC2 based quorum
ecs_mode = "EC2"   # EC2, FARGATE 
ecs_network_mode = "bridge" # bgidge, awsvpc
asg_instance_type = "t2.xlarge"

