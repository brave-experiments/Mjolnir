package terra

var (
	quorum = `
data "aws_security_group" "default" {
  name   = "default"
  vpc_id = "${aws_vpc.this.id}"
}

terraform {
  required_version = ">= 0.10.3" # introduction of Local Values configuration language feature
}

locals {
  max_subnet_length = "${max(length(var.private_subnets), length(var.elasticache_subnets), length(var.database_subnets), length(var.redshift_subnets))}"
  nat_gateway_count = "${var.single_nat_gateway ? 1 : (var.one_nat_gateway_per_az ? length(var.azs) : local.max_subnet_length)}"

  # Use local.vpc_id to give a hint to Terraform that subnets should be deleted before secondary CIDR blocks can be free!
  vpc_id = "${aws_vpc.this.id}"
}

######
# VPC
######
resource "aws_vpc" "this" {
  count = "${var.create_vpc ? 1 : 0}"

  cidr_block                       = "${var.cidr}"
  instance_tenancy                 = "${var.instance_tenancy}"
  enable_dns_hostnames             = "${var.enable_dns_hostnames}"
  enable_dns_support               = "${var.enable_dns_support}"
  assign_generated_ipv6_cidr_block = "${var.assign_generated_ipv6_cidr_block}"

  tags = "${merge(map("Name", format("%s", var.name)), var.tags, var.vpc_tags)}"
}

resource "aws_vpc_ipv4_cidr_block_association" "this" {
  count = "${var.create_vpc && length(var.secondary_cidr_blocks) > 0 ? length(var.secondary_cidr_blocks) : 0}"

  vpc_id = "${aws_vpc.this.id}"

  cidr_block = "${element(var.secondary_cidr_blocks, count.index)}"
}

###################
# DHCP Options Set
###################
resource "aws_vpc_dhcp_options" "this" {
  count = "${var.create_vpc && var.enable_dhcp_options ? 1 : 0}"

  domain_name          = "${var.dhcp_options_domain_name}"
  domain_name_servers  = ["${var.dhcp_options_domain_name_servers}"]
  ntp_servers          = ["${var.dhcp_options_ntp_servers}"]
  netbios_name_servers = ["${var.dhcp_options_netbios_name_servers}"]
  netbios_node_type    = "${var.dhcp_options_netbios_node_type}"

  tags = "${merge(map("Name", format("%s", var.name)), var.tags, var.dhcp_options_tags)}"
}

###############################
# DHCP Options Set Association
###############################
resource "aws_vpc_dhcp_options_association" "this" {
  count = "${var.create_vpc && var.enable_dhcp_options ? 1 : 0}"

  vpc_id          = "${local.vpc_id}"
  dhcp_options_id = "${aws_vpc_dhcp_options.this.id}"
}

###################
# Internet Gateway
###################
resource "aws_internet_gateway" "this" {
  count = "${var.create_vpc && length(var.public_subnets) > 0 ? 1 : 0}"

  vpc_id = "${local.vpc_id}"

  tags = "${merge(map("Name", format("%s", var.name)), var.tags, var.igw_tags)}"
}

################
# Publiс routes
################
resource "aws_route_table" "public" {
  count = "${var.create_vpc && length(var.public_subnets) > 0 ? 1 : 0}"

  vpc_id = "${local.vpc_id}"

  tags = "${merge(map("Name", format("%s-${var.public_subnet_suffix}", var.name)), var.tags, var.public_route_table_tags)}"
}

resource "aws_route" "public_internet_gateway" {
  count = "${var.create_vpc && length(var.public_subnets) > 0 ? 1 : 0}"

  route_table_id         = "${aws_route_table.public.id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.this.id}"

  timeouts {
    create = "5m"
  }
}

#################
# Private routes
# There are as many routing tables as the number of NAT gateways
#################
resource "aws_route_table" "private" {
  count = "${var.create_vpc && local.max_subnet_length > 0 ? local.nat_gateway_count : 0}"

  vpc_id = "${local.vpc_id}"

  tags = "${merge(map("Name", (var.single_nat_gateway ? "${var.name}-${var.private_subnet_suffix}" : format("%s-${var.private_subnet_suffix}-%s", var.name, element(var.azs, count.index)))), var.tags, var.private_route_table_tags)}"

  lifecycle {
    # When attaching VPN gateways it is common to define aws_vpn_gateway_route_propagation
    # resources that manipulate the attributes of the routing table (typically for the private subnets)
    ignore_changes = ["propagating_vgws"]
  }
}


################
# Public subnet
################
resource "aws_subnet" "public" {
  count = "${var.create_vpc && length(var.public_subnets) > 0 && (! var.one_nat_gateway_per_az || length(var.public_subnets) >= length(var.azs)) ? length(var.public_subnets) : 0}"

  vpc_id                  = "${local.vpc_id}"
  cidr_block              = "${element(concat(var.public_subnets, list("")), count.index)}"
  availability_zone       = "${element(var.azs, count.index)}"
  map_public_ip_on_launch = "${var.map_public_ip_on_launch}"

  tags = "${merge(map("Name", format("%s-${var.public_subnet_suffix}-%s", var.name, element(var.azs, count.index))), var.tags, var.public_subnet_tags)}"
}

#################
# Private subnet
#################
resource "aws_subnet" "private" {
  count = "${var.create_vpc && length(var.private_subnets) > 0 ? length(var.private_subnets) : 0}"

  vpc_id            = "${local.vpc_id}"
  cidr_block        = "${var.private_subnets[count.index]}"
  availability_zone = "${element(var.azs, count.index)}"

  tags = "${merge(map("Name", format("%s-${var.private_subnet_suffix}-%s", var.name, element(var.azs, count.index))), var.tags, var.private_subnet_tags)}"
}


##########################
# Route table association
##########################
resource "aws_route_table_association" "private" {
  count = "${var.create_vpc && length(var.private_subnets) > 0 ? length(var.private_subnets) : 0}"

  subnet_id      = "${element(aws_subnet.private.*.id, count.index)}"
  route_table_id = "${element(aws_route_table.private.*.id, (var.single_nat_gateway ? 0 : count.index))}"
}

resource "aws_route_table_association" "public" {
  count = "${var.create_vpc && length(var.public_subnets) > 0 ? length(var.public_subnets) : 0}"

  subnet_id      = "${element(aws_subnet.public.*.id, count.index)}"
  route_table_id = "${aws_route_table.public.id}"
}

###########
# Defaults
###########
resource "aws_default_vpc" "this" {
  count = "${var.manage_default_vpc ? 1 : 0}"

  enable_dns_support   = "${var.default_vpc_enable_dns_support}"
  enable_dns_hostnames = "${var.default_vpc_enable_dns_hostnames}"
  enable_classiclink   = "${var.default_vpc_enable_classiclink}"

  tags = "${merge(map("Name", format("%s", var.default_vpc_name)), var.tags, var.default_vpc_tags)}"
}

variable "create_vpc" {
  description = "Controls if VPC should be created (it affects almost all resources)"
  default     = true
}

variable "name" {
  description = "Name to be used on all the resources as identifier"
  default     = ""
}

variable "cidr" {
  description = "The CIDR block for the VPC. Default value is a valid CIDR, but not acceptable by AWS and should be overridden"
  default     = "10.0.0.0/16"
}

variable "assign_generated_ipv6_cidr_block" {
  description = "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block"
  default     = false
}

variable "secondary_cidr_blocks" {
  description = "List of secondary CIDR blocks to associate with the VPC to extend the IP Address pool"
  default     = []
}

variable "instance_tenancy" {
  description = "A tenancy option for instances launched into the VPC"
  default     = "default"
}

variable "public_subnet_suffix" {
  description = "Suffix to append to public subnets name"
  default     = "public"
}

variable "private_subnet_suffix" {
  description = "Suffix to append to private subnets name"
  default     = "private"
}

variable "intra_subnet_suffix" {
  description = "Suffix to append to intra subnets name"
  default     = "intra"
}

variable "database_subnet_suffix" {
  description = "Suffix to append to database subnets name"
  default     = "db"
}

variable "redshift_subnet_suffix" {
  description = "Suffix to append to redshift subnets name"
  default     = "redshift"
}

variable "elasticache_subnet_suffix" {
  description = "Suffix to append to elasticache subnets name"
  default     = "elasticache"
}

variable "public_subnets" {
  description = "A list of public subnets inside the VPC"
  default     = ["10.0.0.0/24"]
}

variable "private_subnets" {
  description = "A list of private subnets inside the VPC"
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}

variable "database_subnets" {
  description = "A list of database subnets"
  default     = []
}

variable "redshift_subnets" {
  description = "A list of redshift subnets"
  default     = []
}

variable "elasticache_subnets" {
  description = "A list of elasticache subnets"
  default     = []
}

variable "intra_subnets" {
  description = "A list of intra subnets"
  default     = []
}

variable "create_database_subnet_route_table" {
  description = "Controls if separate route table for database should be created"
  default     = false
}

variable "create_redshift_subnet_route_table" {
  description = "Controls if separate route table for redshift should be created"
  default     = false
}

variable "enable_public_redshift" {
  description = "Controls if redshift should have public routing table"
  default     = false
}

variable "create_elasticache_subnet_route_table" {
  description = "Controls if separate route table for elasticache should be created"
  default     = false
}

variable "create_database_subnet_group" {
  description = "Controls if database subnet group should be created"
  default     = true
}

variable "create_elasticache_subnet_group" {
  description = "Controls if elasticache subnet group should be created"
  default     = true
}

variable "create_redshift_subnet_group" {
  description = "Controls if redshift subnet group should be created"
  default     = true
}

variable "create_database_internet_gateway_route" {
  description = "Controls if an internet gateway route for public database access should be created"
  default     = false
}

variable "create_database_nat_gateway_route" {
  description = "Controls if a nat gateway route should be created to give internet access to the database subnets"
  default     = false
}

variable "azs" {
  description = "A list of availability zones in the region"
  default     = ["us-east-2a", "us-east-2b", "us-east-2c"]
}

variable "enable_dns_hostnames" {
  description = "Should be true to enable DNS hostnames in the VPC"
  default     = false
}

variable "enable_dns_support" {
  description = "Should be true to enable DNS support in the VPC"
  default     = true
}

variable "enable_nat_gateway" {
  description = "Should be true if you want to provision NAT Gateways for each of your private networks"
  default     = false
}

variable "single_nat_gateway" {
  description = "Should be true if you want to provision a single shared NAT Gateway across all of your private networks"
  default     = false
}

variable "one_nat_gateway_per_az" {
  description = "Should be true if you want only one NAT Gateway per availability zone. Requires var.azs to be set, and the number of public_subnets created to be greater than or equal to the number of availability zones specified in var.azs."
  default     = false
}

variable "reuse_nat_ips" {
  description = "Should be true if you don't want EIPs to be created for your NAT Gateways and will instead pass them in via the 'external_nat_ip_ids' variable"
  default     = false
}

variable "external_nat_ip_ids" {
  description = "List of EIP IDs to be assigned to the NAT Gateways (used in combination with reuse_nat_ips)"

  default = []
}

variable "enable_dynamodb_endpoint" {
  description = "Should be true if you want to provision a DynamoDB endpoint to the VPC"
  default     = false
}

variable "enable_s3_endpoint" {
  description = "Should be true if you want to provision an S3 endpoint to the VPC"
  default     = false
}

variable "enable_sqs_endpoint" {
  description = "Should be true if you want to provision an SQS endpoint to the VPC"
  default     = false
}

variable "sqs_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for SQS endpoint"
  default     = []
}

variable "sqs_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for SQS endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "sqs_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for SQS endpoint"
  default     = false
}

variable "enable_ssm_endpoint" {
  description = "Should be true if you want to provision an SSM endpoint to the VPC"
  default     = false
}

variable "ssm_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for SSM endpoint"
  default     = []
}

variable "ssm_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for SSM endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ssm_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for SSM endpoint"
  default     = false
}

variable "enable_ssmmessages_endpoint" {
  description = "Should be true if you want to provision a SSMMESSAGES endpoint to the VPC"
  default     = false
}

variable "enable_apigw_endpoint" {
  description = "Should be true if you want to provision an api gateway endpoint to the VPC"
  default     = false
}

variable "apigw_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for API GW  endpoint"
  default     = []
}

variable "apigw_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for API GW endpoint"
  default     = false
}

variable "apigw_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for API GW endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ssmmessages_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for SSMMESSAGES endpoint"
  default     = []
}

variable "ssmmessages_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for SSMMESSAGES endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ssmmessages_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for SSMMESSAGES endpoint"
  default     = false
}

variable "enable_ec2_endpoint" {
  description = "Should be true if you want to provision an EC2 endpoint to the VPC"
  default     = false
}

variable "ec2_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for EC2 endpoint"
  default     = []
}

variable "ec2_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for EC2 endpoint"
  default     = false
}

variable "ec2_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for EC2 endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "enable_ec2messages_endpoint" {
  description = "Should be true if you want to provision an EC2MESSAGES endpoint to the VPC"
  default     = false
}

variable "ec2messages_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for EC2MESSAGES endpoint"
  default     = []
}

variable "ec2messages_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for EC2MESSAGES endpoint"
  default     = false
}

variable "ec2messages_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for EC2MESSAGES endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "enable_ecr_api_endpoint" {
  description = "Should be true if you want to provision an ecr api endpoint to the VPC"
  default     = false
}

variable "ecr_api_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for ECR api endpoint. If omitted, private subnets will be used."
  default     = []
}

variable "ecr_api_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for ECR API endpoint"
  default     = false
}

variable "ecr_api_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for ECR API endpoint"
  default     = []
}

variable "enable_ecr_dkr_endpoint" {
  description = "Should be true if you want to provision an ecr dkr endpoint to the VPC"
  default     = false
}

variable "ecr_dkr_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for ECR dkr endpoint. If omitted, private subnets will be used."
  default     = []
}

variable "ecr_dkr_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for ECR DKR endpoint"
  default     = false
}

variable "ecr_dkr_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for ECR DKR endpoint"
  default     = []
}

variable "enable_kms_endpoint" {
  description = "Should be true if you want to provision a KMS endpoint to the VPC"
  default     = false
}

variable "kms_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for KMS endpoint"
  default     = []
}

variable "kms_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for KMS endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "kms_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for KMS endpoint"
  default     = false
}

variable "enable_ecs_endpoint" {
  description = "Should be true if you want to provision a ECS endpoint to the VPC"
  default     = false
}

variable "ecs_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for ECS endpoint"
  default     = []
}

variable "ecs_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for ECS endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ecs_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for ECS endpoint"
  default     = false
}

variable "enable_ecs_agent_endpoint" {
  description = "Should be true if you want to provision a ECS Agent endpoint to the VPC"
  default     = false
}

variable "ecs_agent_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for ECS Agent endpoint"
  default     = []
}

variable "ecs_agent_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for ECS Agent endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ecs_agent_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for ECS Agent endpoint"
  default     = false
}

variable "enable_ecs_telemetry_endpoint" {
  description = "Should be true if you want to provision a ECS Telemetry endpoint to the VPC"
  default     = false
}

variable "ecs_telemetry_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for ECS Telemetry endpoint"
  default     = []
}

variable "ecs_telemetry_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for ECS Telemetry endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "ecs_telemetry_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for ECS Telemetry endpoint"
  default     = false
}

variable "enable_logs_endpoint" {
  description = "Should be true if you want to provision a CloudWatch Logs endpoint to the VPC"
  default     = false
}

variable "logs_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for CloudWatch Logs endpoint"
  default     = []
}

variable "logs_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for CloudWatch Logs endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "logs_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for CloudWatch Logs endpoint"
  default     = false
}

variable "enable_cloudtrail_endpoint" {
  description = "Should be true if you want to provision a CloudTrail endpoint to the VPC"
  default     = false
}

variable "cloudtrail_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for CloudTrail endpoint"
  default     = []
}

variable "cloudtrail_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for CloudTrail endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "cloudtrail_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for CloudTrail endpoint"
  default     = false
}

variable "enable_elasticloadbalancing_endpoint" {
  description = "Should be true if you want to provision a Elastic Load Balancing endpoint to the VPC"
  default     = false
}

variable "elasticloadbalancing_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for Elastic Load Balancing endpoint"
  default     = []
}

variable "elasticloadbalancing_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for Elastic Load Balancing endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "elasticloadbalancing_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for Elastic Load Balancing endpoint"
  default     = false
}

variable "enable_sns_endpoint" {
  description = "Should be true if you want to provision a SNS endpoint to the VPC"
  default     = false
}

variable "sns_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for SNS endpoint"
  default     = []
}

variable "sns_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for SNS endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "sns_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for SNS endpoint"
  default     = false
}

variable "enable_events_endpoint" {
  description = "Should be true if you want to provision a CloudWatch Events endpoint to the VPC"
  default     = false
}

variable "events_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for CloudWatch Events endpoint"
  default     = []
}

variable "events_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for CloudWatch Events endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "events_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for CloudWatch Events endpoint"
  default     = false
}

variable "enable_monitoring_endpoint" {
  description = "Should be true if you want to provision a CloudWatch Monitoring endpoint to the VPC"
  default     = false
}

variable "monitoring_endpoint_security_group_ids" {
  description = "The ID of one or more security groups to associate with the network interface for CloudWatch Monitoring endpoint"
  default     = []
}

variable "monitoring_endpoint_subnet_ids" {
  description = "The ID of one or more subnets in which to create a network interface for CloudWatch Monitoring endpoint. Only a single subnet within an AZ is supported. If omitted, private subnets will be used."
  default     = []
}

variable "monitoring_endpoint_private_dns_enabled" {
  description = "Whether or not to associate a private hosted zone with the specified VPC for CloudWatch Monitoring endpoint"
  default     = false
}

variable "map_public_ip_on_launch" {
  description = "Should be false if you do not want to auto-assign public IP on launch"
  default     = true
}

variable "enable_vpn_gateway" {
  description = "Should be true if you want to create a new VPN Gateway resource and attach it to the VPC"
  default     = false
}

variable "vpn_gateway_id" {
  description = "ID of VPN Gateway to attach to the VPC"
  default     = ""
}

variable "amazon_side_asn" {
  description = "The Autonomous System Number (ASN) for the Amazon side of the gateway. By default the virtual private gateway is created with the current default Amazon ASN."
  default     = "64512"
}

variable "propagate_private_route_tables_vgw" {
  description = "Should be true if you want route table propagation"
  default     = false
}

variable "propagate_public_route_tables_vgw" {
  description = "Should be true if you want route table propagation"
  default     = false
}

variable "tags" {
  description = "A map of tags to add to all resources"
  default     = {}
}

variable "vpc_tags" {
  description = "Additional tags for the VPC"
  default     = {}
}

variable "igw_tags" {
  description = "Additional tags for the internet gateway"
  default     = {}
}

variable "public_subnet_tags" {
  description = "Additional tags for the public subnets"
  default     = {}
}

variable "private_subnet_tags" {
  description = "Additional tags for the private subnets"
  default     = {}
}

variable "public_route_table_tags" {
  description = "Additional tags for the public route tables"
  default     = {}
}

variable "private_route_table_tags" {
  description = "Additional tags for the private route tables"
  default     = {}
}

variable "database_route_table_tags" {
  description = "Additional tags for the database route tables"
  default     = {}
}

variable "redshift_route_table_tags" {
  description = "Additional tags for the redshift route tables"
  default     = {}
}

variable "elasticache_route_table_tags" {
  description = "Additional tags for the elasticache route tables"
  default     = {}
}

variable "intra_route_table_tags" {
  description = "Additional tags for the intra route tables"
  default     = {}
}

variable "database_subnet_tags" {
  description = "Additional tags for the database subnets"
  default     = {}
}

variable "database_subnet_group_tags" {
  description = "Additional tags for the database subnet group"
  default     = {}
}

variable "redshift_subnet_tags" {
  description = "Additional tags for the redshift subnets"
  default     = {}
}

variable "redshift_subnet_group_tags" {
  description = "Additional tags for the redshift subnet group"
  default     = {}
}

variable "elasticache_subnet_tags" {
  description = "Additional tags for the elasticache subnets"
  default     = {}
}

variable "intra_subnet_tags" {
  description = "Additional tags for the intra subnets"
  default     = {}
}

variable "public_acl_tags" {
  description = "Additional tags for the public subnets network ACL"
  default     = {}
}

variable "private_acl_tags" {
  description = "Additional tags for the private subnets network ACL"
  default     = {}
}

variable "intra_acl_tags" {
  description = "Additional tags for the intra subnets network ACL"
  default     = {}
}

variable "database_acl_tags" {
  description = "Additional tags for the database subnets network ACL"
  default     = {}
}

variable "redshift_acl_tags" {
  description = "Additional tags for the redshift subnets network ACL"
  default     = {}
}

variable "elasticache_acl_tags" {
  description = "Additional tags for the elasticache subnets network ACL"
  default     = {}
}

variable "dhcp_options_tags" {
  description = "Additional tags for the DHCP option set (requires enable_dhcp_options set to true)"
  default     = {}
}

variable "nat_gateway_tags" {
  description = "Additional tags for the NAT gateways"
  default     = {}
}

variable "nat_eip_tags" {
  description = "Additional tags for the NAT EIP"
  default     = {}
}

variable "vpn_gateway_tags" {
  description = "Additional tags for the VPN gateway"
  default     = {}
}

variable "enable_dhcp_options" {
  description = "Should be true if you want to specify a DHCP options set with a custom domain name, DNS servers, NTP servers, netbios servers, and/or netbios server type"
  default     = false
}

variable "dhcp_options_domain_name" {
  description = "Specifies DNS name for DHCP options set (requires enable_dhcp_options set to true)"
  default     = ""
}

variable "dhcp_options_domain_name_servers" {
  description = "Specify a list of DNS server addresses for DHCP options set, default to AWS provided (requires enable_dhcp_options set to true)"

  default = ["AmazonProvidedDNS"]
}

variable "dhcp_options_ntp_servers" {
  description = "Specify a list of NTP servers for DHCP options set (requires enable_dhcp_options set to true)"

  default = []
}

variable "dhcp_options_netbios_name_servers" {
  description = "Specify a list of netbios servers for DHCP options set (requires enable_dhcp_options set to true)"

  default = []
}

variable "dhcp_options_netbios_node_type" {
  description = "Specify netbios node_type for DHCP options set (requires enable_dhcp_options set to true)"
  default     = ""
}

variable "manage_default_vpc" {
  description = "Should be true to adopt and manage Default VPC"
  default     = false
}

variable "default_vpc_name" {
  description = "Name to be used on the Default VPC"
  default     = ""
}

variable "default_vpc_enable_dns_support" {
  description = "Should be true to enable DNS support in the Default VPC"
  default     = true
}

variable "default_vpc_enable_dns_hostnames" {
  description = "Should be true to enable DNS hostnames in the Default VPC"
  default     = false
}

variable "default_vpc_enable_classiclink" {
  description = "Should be true to enable ClassicLink in the Default VPC"
  default     = false
}

variable "default_vpc_tags" {
  description = "Additional tags for the Default VPC"
  default     = {}
}

variable "manage_default_network_acl" {
  description = "Should be true to adopt and manage Default Network ACL"
  default     = false
}

variable "default_network_acl_name" {
  description = "Name to be used on the Default Network ACL"
  default     = ""
}

variable "default_network_acl_tags" {
  description = "Additional tags for the Default Network ACL"
  default     = {}
}

variable "public_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for public subnets"
  default     = false
}

variable "private_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for private subnets"
  default     = false
}

variable "intra_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for intra subnets"
  default     = false
}

variable "database_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for database subnets"
  default     = false
}

variable "redshift_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for redshift subnets"
  default     = false
}

variable "elasticache_dedicated_network_acl" {
  description = "Whether to use dedicated network ACL (not default) and custom rules for elasticache subnets"
  default     = false
}

variable "default_network_acl_ingress" {
  description = "List of maps of ingress rules to set on the Default Network ACL"

  default = [{
    rule_no    = 100
    action     = "allow"
    from_port  = 0
    to_port    = 0
    protocol   = "-1"
    cidr_block = "0.0.0.0/0"
  },
    {
      rule_no         = 101
      action          = "allow"
      from_port       = 0
      to_port         = 0
      protocol        = "-1"
      ipv6_cidr_block = "::/0"
    },
  ]
}

variable "default_network_acl_egress" {
  description = "List of maps of egress rules to set on the Default Network ACL"

  default = [{
    rule_no    = 100
    action     = "allow"
    from_port  = 0
    to_port    = 0
    protocol   = "-1"
    cidr_block = "0.0.0.0/0"
  },
    {
      rule_no         = 101
      action          = "allow"
      from_port       = 0
      to_port         = 0
      protocol        = "-1"
      ipv6_cidr_block = "::/0"
    },
  ]
}

variable "public_inbound_acl_rules" {
  description = "Public subnets inbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "public_outbound_acl_rules" {
  description = "Public subnets outbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "private_inbound_acl_rules" {
  description = "Private subnets inbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "private_outbound_acl_rules" {
  description = "Private subnets outbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "intra_inbound_acl_rules" {
  description = "Intra subnets inbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "intra_outbound_acl_rules" {
  description = "Intra subnets outbound network ACLs"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "database_inbound_acl_rules" {
  description = "Database subnets inbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "database_outbound_acl_rules" {
  description = "Database subnets outbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "redshift_inbound_acl_rules" {
  description = "Redshift subnets inbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "redshift_outbound_acl_rules" {
  description = "Redshift subnets outbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "elasticache_inbound_acl_rules" {
  description = "Elasticache subnets inbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

variable "elasticache_outbound_acl_rules" {
  description = "Elasticache subnets outbound network ACL rules"

  default = [
    {
      rule_number = 100
      rule_action = "allow"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_block  = "0.0.0.0/0"
    },
  ]
}

resource "aws_iam_role" "bastion" {
  name = "${local.default_bastion_resource_name}"
  path = "/"

  assume_role_policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": [
          "ec2.amazonaws.com"
        ]
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "bastion" {
  statement {
    sid = "AllowS3"

    actions = [
      "s3:*",
    ]

    resources = [
      "arn:aws:s3:::${local.quorum_bucket}",
      "arn:aws:s3:::${local.quorum_bucket}/*",
    ]
  }

  statement {
    sid = "AllowS3Bastion"

    actions = [
      "s3:*",
    ]

    resources = [
      "arn:aws:s3:::${local.bastion_bucket}",
      "arn:aws:s3:::${local.bastion_bucket}/*",
    ]
  }

  statement {
    sid = "AllowKMSAccess"

    actions = [
      "kms:*",
    ]

    resources = [
      "${aws_kms_key.bucket.arn}",
    ]
  }

  statement {
    sid = "AllowECS"

    actions = [
      "ecs:*",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    sid = "AllowEC2"

    actions = [
      "ec2:*",
    ]

    resources = [
      "*",
    ]
  }
}

resource "aws_iam_instance_profile" "bastion" {
  name = "${local.default_bastion_resource_name}"
  role = "${aws_iam_role.bastion.name}"
}

resource "aws_iam_policy" "bastion" {
  name        = "quorum-bastion-policy-${var.network_name}"
  path        = "/"
  description = "This policy allows task to access S3 bucket and ECS"
  policy      = "${data.aws_iam_policy_document.bastion.json}"
}

resource "aws_iam_role_policy_attachment" "bastion" {
  role       = "${aws_iam_role.bastion.id}"
  policy_arn = "${aws_iam_policy.bastion.arn}"
}

locals {
  default_bastion_resource_name = "${format("quorum-bastion-%s", var.network_name)}"
  ethstats_docker_image         = "puppeth/ethstats:latest"
  ethstats_port                 = 3000
  bastion_bucket    = "${var.region}-bastion-${lower(var.network_name)}-${random_id.bucket_postfix.hex}"
}

data "aws_ami" "this" {
  most_recent = true

  filter {
    name = "name"

    values = [
      "amzn2-ami-hvm-*",
    ]
  }

  filter {
    name = "virtualization-type"

    values = [
      "hvm",
    ]
  }

  filter {
    name = "architecture"

    values = [
      "x86_64",
    ]
  }

  owners = [
    "137112412989",
  ]

  # amazon
}

resource "random_id" "ethstat_secret" {
  byte_length = 16
}

resource "tls_private_key" "ssh" {
  algorithm = "RSA"
  rsa_bits  = "2048"
}

resource "aws_key_pair" "ssh" {
  public_key = "${tls_private_key.ssh.public_key_openssh}"
  key_name   = "${local.default_bastion_resource_name}"
}

resource "local_file" "private_key" {
  filename = "${path.module}/quorum-${var.network_name}.pem"
  content  = "${tls_private_key.ssh.private_key_pem}"

  provisioner "local-exec" {
    on_failure = "continue"
    command    = "chmod 600 ${self.filename}"
  }
}

resource "aws_instance" "bastion" {
  ami           = "${data.aws_ami.this.id}"
  instance_type = "t2.large"

  vpc_security_group_ids = [
    "${aws_security_group.quorum.id}",
    "${aws_security_group.bastion-ssh.id}",
    "${aws_security_group.bastion-ethstats.id}",
  ]

  subnet_id                   = "${aws_subnet.public.id}"
  associate_public_ip_address = "true"
  key_name                    = "${aws_key_pair.ssh.key_name}"
  iam_instance_profile        = "${aws_iam_instance_profile.bastion.name}"

  user_data = <<EOF
#!/bin/bash

set -e

# START: added per suggestion from AWS support to mitigate an intermittent failures from yum update
sleep 20
yum clean all
yum repolist
# END

yum -y update
yum -y install jq
amazon-linux-extras install docker -y
systemctl enable docker
systemctl start docker

curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker pull ${local.quorum_docker_image}
docker pull prom/prometheus
docker pull prom/node-exporter:latest
docker pull grafana/grafana:latest
mkdir -p /opt/prometheus
docker run -d -e "WS_SECRET=${random_id.ethstat_secret.hex}" -p ${local.ethstats_port}:${local.ethstats_port} ${local.ethstats_docker_image}
EOF

  provisioner "remote-exec" {
    inline = [
      "sudo yum -y update",
      "sudo yum -y install jq",
      "sudo amazon-linux-extras install docker -y",
      "sudo systemctl enable docker",
      "sudo systemctl start docker",
      "printf 'FROM alpine\nCOPY --from=trajano/alpine-libfaketime  /faketime.so /lib/faketime.so\n' > /tmp/Dockerfile.libfaketime",
      "sudo docker build -f /tmp/Dockerfile.libfaketime . -t libfaketime:latest",
      "sudo docker run -v $PWD:/tmp --rm --entrypoint cp libfaketime:latest /lib/faketime.so /tmp/libfaketime.so",
      "sudo aws s3 cp libfaketime.so s3://${local.bastion_bucket}/libs/libfaketime.so"
    ]

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "ec2-user"
      private_key = "${tls_private_key.ssh.private_key_pem}"
      timeout     = "10m"
    }
  }


  tags = "${merge(local.common_tags, map("Name", local.default_bastion_resource_name))}"
}

resource "local_file" "bootstrap" {
  filename = "${path.module}/generated-bootstrap.sh"

  content = <<EOF
#!/bin/bash

set -e

export AWS_DEFAULT_REGION=${var.region}
export TASK_REVISION=${aws_ecs_task_definition.quorum.revision}
sudo rm -rf ${local.shared_volume_container_path}
sudo mkdir -p ${local.shared_volume_container_path}/mappings
sudo mkdir -p ${local.privacy_addresses_folder}

# Faketime array ( ClockSkew )
old_IFS=$IFS
IFS=',' faketime=(${join(" ", var.faketime)})
IFS=$${old_IFS}
counter="$${#faketime[@]}"

while [ $counter -gt 0 ]
do
    echo -n "$${faketime[-1]}" > ./$counter
    faketime=($${faketime[@]::$counter})
    sudo aws s3 cp ./$counter s3://${local.bastion_bucket}/clockSkew/
    counter=$((counter - 1))
done

count=0
while [ $count -lt ${var.number_of_nodes} ]
do
  count=$(ls ${local.privacy_addresses_folder} | grep ^ip | wc -l)
  sudo aws s3 cp --recursive s3://${local.s3_revision_folder}/ ${local.shared_volume_container_path}/ > /dev/null 2>&1 \
    | echo Wait for nodes in Quorum network being up ... $count/${var.number_of_nodes}
  sleep 1
done

if which jq >/dev/null; then
  echo "Found jq"
else
  echo "jq not found. Instaling ..."
  sudo yum -y install jq
fi

for t in aws ecs list-tasks --cluster ${local.ecs_cluster_name} | jq -r .taskArns[]
do
  task_metadata=$(aws ecs describe-tasks --cluster ${local.ecs_cluster_name} --tasks $t)
  HOST_IP=$(echo $task_metadata | jq -r '.tasks[0] | .containers[] | select(.name == "${local.quorum_run_container_name}") | .networkInterfaces[] | .privateIpv4Address')
  if [ "${var.ecs_mode}" == "EC2" ]
  then
    CONTAINER_INSTANCE_ARN=$(aws ecs describe-tasks --tasks $t --cluster ${local.ecs_cluster_name} | jq -r '.tasks[] | .containerInstanceArn')
    EC2_INSTANCE_ID=$(aws ecs  describe-container-instances --container-instances $CONTAINER_INSTANCE_ARN --cluster ${local.ecs_cluster_name} |jq -r '.containerInstances[] | .ec2InstanceId')
    HOST_IP=$(aws ec2 describe-instances --instance-ids $EC2_INSTANCE_ID | jq -r '.Reservations[0] | .Instances[] | .PublicIpAddress')
  fi
  group=$(echo $task_metadata | jq -r '.tasks[0] | .group')
  taskArn=$(echo $task_metadata | jq -r '.tasks[0] | .taskDefinitionArn')
  # only care about new task
  if [[ "$taskArn" == *:$TASK_REVISION ]]; then
     echo $group | sudo tee ${local.shared_volume_container_path}/mappings/${local.normalized_host_ip}
  fi
done

cat <<SS | sudo tee ${local.shared_volume_container_path}/quorum_metadata
quorum:
  nodes:
SS
nodes=(${join(" ", aws_ecs_service.quorum.*.name)})
cd ${local.shared_volume_container_path}/mappings
for idx in "$${!nodes[@]}"
do
  f=$(grep -l $${nodes[$idx]} *)
  ip=$(cat ${local.hosts_folder}/$f)
  nodeIdx=$((idx+1))
  script="/usr/local/bin/Node$nodeIdx"
  cat <<SS | sudo tee $script
#!/bin/bash

sudo docker run --rm -it ${local.quorum_docker_image} attach http://$ip:${local.quorum_rpc_port} $@
SS
  sudo chmod +x $script
  cat <<SS | sudo tee -a ${local.shared_volume_container_path}/quorum_metadata
    Node$nodeIdx:
      privacy-address: $(cat ${local.privacy_addresses_folder}/$f)
      url: http://$ip:${local.quorum_rpc_port}
      third-party-url: http://$ip:${local.tessera_thirdparty_port}
SS
done

cat <<SS | sudo tee /opt/prometheus/prometheus.yml
global:
  scrape_interval:     15s # By default, scrape targets every 15 seconds.

  # Attach these labels to any time series or alerts when communicating with
  # external systems (federation, remote storage, Alertmanager).
  external_labels:
    monitor: 'monitor'

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
- job_name: 'node'
  static_configs:
  - targets: ['node-exporter:9100','gethexporter:9090']
  file_sd_configs:
  - files:
    - 'targets.json'
SS

cat <<SS | sudo tee /opt/prometheus/docker-compose.yml
# docker-compose.yml
version: '2'
services:
    prometheus:
        image: prom/prometheus:latest
        volumes:
            - /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
            - /opt/prometheus/targets.json:/etc/prometheus/targets.json
        command:
            - '--config.file=/etc/prometheus/prometheus.yml'
        ports:
            - '9090:9090'
    node-exporter:
        image: prom/node-exporter:latest
        ports:
            - '9100:9100'
    grafana:
        image: grafana/grafana:latest
        environment:
            - GF_SECURITY_ADMIN_PASSWORD=my-pass
        depends_on:
            - prometheus
        ports:
            - '3001:3000'
    gethexporter:
        image: hunterlong/gethexporter:latest
        environment:
            - GETH=http://mygethserverhere:22000
        depends_on:
            - prometheus
        ports:
            - '9191:9090'
SS

count=$(ls ${local.privacy_addresses_folder} | grep ^ip | wc -l)
target_file=/tmp/targets.json
i=0
echo '[' > $target_file
for idx in "$${!nodes[@]}"
do
  f=$(grep -l $${nodes[$idx]} *)
  ip=$(cat ${local.hosts_folder}/$f)
  i=$(($i+1))
  if [ $i -lt "$count" ]; then
    echo '{ "targets": ["'$ip':9100"] },' >> $target_file
  else
    echo '{ "targets": ["'$ip':9100"] }'  >> $target_file
  fi
done
echo ']' >> $target_file
sudo mv $target_file /opt/prometheus/
sudo sed -i s"/mygethserverhere/$ip/" /opt/prometheus/docker-compose.yml
sudo /usr/local/bin/docker-compose -f /opt/prometheus/docker-compose.yml up -d --force-recreate
EOF
}

resource "null_resource" "bastion_remote_exec" {
  triggers {
    bastion             = "${aws_instance.bastion.public_dns}"
    ecs_task_definition = "${aws_ecs_task_definition.quorum.revision}"
    script              = "${md5(local_file.bootstrap.content)}"
  }

  provisioner "remote-exec" {
    script = "${local_file.bootstrap.filename}"

    connection {
      host        = "${aws_instance.bastion.public_ip}"
      user        = "ec2-user"
      private_key = "${tls_private_key.ssh.private_key_pem}"
      timeout     = "10m"
    }
  }
}

resource "aws_s3_bucket" "bastion" {
  bucket        = "${local.bastion_bucket}"
  region        = "${var.region}"
  force_destroy = true

  versioning {
    enabled = true
  }
}
resource "aws_security_group" "bastion-ssh" {
  vpc_id      = "${local.vpc_id}"
  name        = "quorum-bastion-ssh-${var.network_name}"
  description = "Security group used by Bastion node to access Quorum network ${var.network_name}"

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

  tags = "${merge(local.common_tags, map("Name", format("quorum-bastion-ssh-%s", var.network_name)))}"
}

resource "aws_security_group" "bastion-ethstats" {
  vpc_id      = "${local.vpc_id}"
  name        = "quorum-bastion-ethstats-${var.network_name}"
  description = "Security group used by external to access ethstats for Quorum network ${var.network_name}"

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

  tags = "${merge(local.common_tags, map("Name", format("quorum-bastion-ethstats-%s", var.network_name)))}"
}

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

/*
 * Determine most recent ECS optimized AMI
 */
data "aws_ami" "ecs_ami" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-*-amazon-ecs-optimized"]
  }
}


/*
 * Create ECS IAM Instance Role and Policy
 */
resource "random_id" "code" {
  byte_length = 4
}

resource "aws_iam_role" "ecsInstanceRole" {
  name               = "ecsInstanceRole-${random_id.code.hex}"
  assume_role_policy = "${var.ecsInstanceRoleAssumeRolePolicy}"
}

resource "aws_iam_role_policy" "ecsInstanceRolePolicy" {
  name   = "ecsInstanceRolePolicy-${random_id.code.hex}"
  role   = "${aws_iam_role.ecsInstanceRole.id}"
  policy = "${var.ecsInstancerolePolicy}"
}

/*
 * Create ECS IAM Service Role and Policy
 */
resource "aws_iam_role" "ecsServiceRole" {
  name               = "ecsServiceRole-${random_id.code.hex}"
  assume_role_policy = "${var.ecsServiceRoleAssumeRolePolicy}"
}

resource "aws_iam_role_policy" "ecsServiceRolePolicy" {
  name   = "ecsServiceRolePolicy-${random_id.code.hex}"
  role   = "${aws_iam_role.ecsServiceRole.id}"
  policy = "${var.ecsServiceRolePolicy}"
}

resource "aws_iam_instance_profile" "ecsInstanceProfile" {
  name = "ecsInstanceProfile-${random_id.code.hex}"
  role = "${aws_iam_role.ecsInstanceRole.name}"
}

/*
 * ECS related variables
 */


// Optional:

variable "ecsInstanceRoleAssumeRolePolicy" {
  type = "string"

  default = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

variable "ecsInstancerolePolicy" {
  type = "string"

  default = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ecs:CreateCluster",
        "ecs:DeregisterContainerInstance",
        "ecs:DiscoverPollEndpoint",
        "ecs:Poll",
        "ecs:RegisterContainerInstance",
        "ecs:StartTelemetrySession",
        "ecs:Submit*",
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

variable "ecsServiceRoleAssumeRolePolicy" {
  type = "string"

  default = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

variable "ecsServiceRolePolicy" {
  default = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:Describe*",
        "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
        "elasticloadbalancing:DeregisterTargets",
        "elasticloadbalancing:Describe*",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:RegisterTargets"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

locals {
  host_ip_file         = "${local.shared_volume_container_path}/host_ip"
  task_revision_file   = "${local.shared_volume_container_path}/task_revision"
  service_file         = "${local.shared_volume_container_path}/service"
  account_address_file = "${local.shared_volume_container_path}/first_account_address"
  hosts_folder         = "${local.shared_volume_container_path}/hosts"
  libfaketime_folder  =  "${local.shared_volume_container_path}/lib"

  metadata_bootstrap_container_status_file = "${local.shared_volume_container_path}/metadata_bootstrap_container_status"

  // For S3 related operations
  s3_revision_folder = "${local.quorum_bucket}/rev_$TASK_REVISION"
  s3_libfaketime_file = "${local.bastion_bucket}/libs/libfaketime.so"
  normalized_host_ip = "ip_$(echo $HOST_IP | sed -e 's/\\./_/g')"

  node_key_bootstrap_commands = [
    "mkdir -p ${local.quorum_data_dir}/geth",
    "echo \"\" > ${local.quorum_password_file}",
    "bootnode -genkey ${local.quorum_data_dir}/geth/nodekey",
    "export NODE_ID=$(bootnode -nodekey ${local.quorum_data_dir}/geth/nodekey -writeaddress)",
    "echo Creating an account for this node",
    "geth --datadir ${local.quorum_data_dir} account new --password ${local.quorum_password_file}",
    "export KEYSTORE_FILE=$(ls ${local.quorum_data_dir}/keystore/ | head -n1)",
    "export ACCOUNT_ADDRESS=$(cat ${local.quorum_data_dir}/keystore/$KEYSTORE_FILE | sed 's/^.*\"address\":\"\\([^\"]*\\)\".*$/\\1/g')",
    "echo Writing account address $ACCOUNT_ADDRESS to ${local.account_address_file}",
    "echo $ACCOUNT_ADDRESS > ${local.account_address_file}",
    "echo Writing Node Id [$NODE_ID] to ${local.node_id_file}",
    "echo $NODE_ID > ${local.node_id_file}",
  ]

  node_key_bootstrap_container_definition = {
    name      = "${local.node_key_bootstrap_container_name}"
    image     = "${local.quorum_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
      }
    }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    environments = []

    portMappings = []

    volumesFrom = []

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -f ${local.node_id_file} ];",
      ]
    }

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.node_key_bootstrap_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }

  // this is very BADDDDDD but for now i don't have any other better option
  validator_address_program = <<EOP
package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing enode value")
		os.Exit(1)
	}
	enode := os.Args[1]
	nodeId, err := discover.HexID(enode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pub, err := nodeId.Pubkey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("0x%s\n", hex.EncodeToString(crypto.PubkeyToAddress(*pub).Bytes()))
}
EOP

  // bootstrap the extraData, this must be used inside metadata_bootstrap_commands to inherit metadata info
  istanbul_bootstrap_commands = [
    "apk add --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/community go=1.9.4-r0",
    "apk add git gcc musl-dev linux-headers",
    "git clone ${element(local.consensus_config_map["git_url"], 0)} /istanbul-tools/src/github.com/getamis/istanbul-tools",
    "export GOPATH=/istanbul-tools",
    "export GOROOT=/usr/lib/go",
    "echo '${local.validator_address_program}' > /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go",
    "all=\"\"; for f in ls ${local.node_ids_folder}; do address=$(cat ${local.node_ids_folder}/$f); all=\"$all,$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go $address)\"; done",
    "all=\"$${all:1}\"",
    "echo Validator Addresses: $all",
    "extraData=\"\\\"$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/cmd/istanbul/main.go extra encode --validators $all | awk -F: '{print $2}' | tr -d ' ')\\\"\"",
  ]

  metadata_bootstrap_commands = [
    "set -e",
    "echo Wait until Node Key is ready ...",
    "while [ ! -f \"${local.node_id_file}\" ]; do sleep 1; done",
    "apk update",
    "apk add curl jq",
    "export TASK_REVISION=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq '.Revision' -r)",
    "echo \"Task Revision: $TASK_REVISION\"",
    "echo $TASK_REVISION > ${local.task_revision_file}",

    //"export HOST_IP=$(curl -s $ECS_CONTAINER_METADATA_URI/task |  jq '.Containers[] | select(.Name== \"${local.metadata_bootstrap_container_name}\") | .Networks[] | select(.NetworkMode== \"bridge\") | .IPv4Addresses[0]' -r)",
    //"export HOST_IP=$(/sbin/ip route|awk '/default/ { print $3 }')",
    "export HOST_IP=$(/usr/bin/curl http://169.254.169.254/latest/meta-data/public-ipv4)",

    "echo \"Host IP: $HOST_IP\"",
    "echo $HOST_IP > ${local.host_ip_file}",
    "export TASK_ARN=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')",
    "export REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')",
    "aws ecs describe-tasks --region $REGION --cluster ${local.ecs_cluster_name} --tasks $TASK_ARN | jq -r '.tasks[0] | .group' > ${local.service_file}",
    "mkdir -p ${local.hosts_folder}",
    "mkdir -p ${local.node_ids_folder}",
    "mkdir -p ${local.accounts_folder}",
    "mkdir -p ${local.libfaketime_folder}",
    "aws s3 cp s3://${local.s3_libfaketime_file} ${local.libfaketime_folder}/libfaketime.so",
    "aws s3 cp ${local.node_id_file} s3://${local.s3_revision_folder}/nodeids/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "aws s3 cp ${local.host_ip_file} s3://${local.s3_revision_folder}/hosts/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
    "aws s3 cp ${local.account_address_file} s3://${local.s3_revision_folder}/accounts/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",

    // Gather all IPs
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.hosts_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/hosts ${local.hosts_folder} > /dev/null 2>&1 | echo \"Wait for other containers to report their IPs ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All containers have reported their IPs\"",

    // Gather all Accounts
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.accounts_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/accounts ${local.accounts_folder} > /dev/null 2>&1 | echo \"Wait for other nodes to report their accounts ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All nodes have registered accounts\"",

    // Gather all Node IDs
    "count=0; while [ $count -lt ${var.number_of_nodes} ]; do count=$(ls ${local.node_ids_folder} | grep ^ip | wc -l); aws s3 cp --recursive s3://${local.s3_revision_folder}/nodeids ${local.node_ids_folder} > /dev/null 2>&1 | echo \"Wait for other nodes to report their IDs ... $count/${var.number_of_nodes}\"; sleep 1; done",

    "echo \"All nodes have registered their IDs\"",

    // Prepare Genesis file
    "alloc=\"\"; for f in ls ${local.accounts_folder}; do address=$(cat ${local.accounts_folder}/$f); alloc=\"$alloc,\\\"$address\\\": { \"balance\": \"\\\"1000000000000000000000000000\\\"\"}\"; done",

    "alloc=\"{$${alloc:1}}\"",
    "extraData=\"\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\"\"",
    "${var.consensus_mechanism == "istanbul" ? join("\n", local.istanbul_bootstrap_commands) : ""}",
    "mixHash=\"\\\"${element(local.consensus_config_map["genesis_mixHash"], 0)}\\\"\"",
    "difficulty=\"\\\"${element(local.consensus_config_map["genesis_difficulty"], 0)}\\\"\"",
    "echo '${replace(jsonencode(local.genesis), "/\"(true|false|[0-9]+)\"/", "$1")}' | jq \". + { alloc : $alloc, extraData: $extraData, mixHash: $mixHash, difficulty: $difficulty}${var.consensus_mechanism == "istanbul" ? " | .config=.config + {istanbul: {epoch: 30000, policy: 0} }" : ""}\" > ${local.genesis_file}",
    "cat ${local.genesis_file}",

    // Write status
    "echo \"Done!\" > ${local.metadata_bootstrap_container_status_file}",

    "echo Wait until privacy engine initialized ...",
    "while [ ! -f \"${local.tx_privacy_engine_address_file}\" ]; do sleep 1; done",
    "aws s3 cp ${local.tx_privacy_engine_address_file} s3://${local.s3_revision_folder}/privacyaddresses/${local.normalized_host_ip} --sse aws:kms --sse-kms-key-id ${aws_kms_key.bucket.arn}",
  ]

  metadata_bootstrap_container_definition = {
    name      = "${local.metadata_bootstrap_container_name}"
    image     = "${local.aws_cli_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
      }
    }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    environments = []

    portMappings = []

    volumesFrom = [
      {
        sourceContainer = "${local.node_key_bootstrap_container_name}"
      },
    ]

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -f ${local.metadata_bootstrap_container_status_file} ];",
      ]
    }

    entryPoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.metadata_bootstrap_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
}

locals {
  shared_volume_name             = "quorum_shared_volume"
  shared_volume_container_path   = "/qdata"
  tx_privacy_engine_socket_file  = "${local.shared_volume_container_path}/tm.ipc"
  tx_privacy_engine_address_file = "${element(compact(local.tx_privacy_address_files), 0)}"

  node_key_bootstrap_container_name           = "node-key-bootstrap"
  metadata_bootstrap_container_name           = "metamain-bootstrap"
  quorum_run_container_name                   = "quorum-run"
  tx_privacy_engine_run_container_name        = "${var.tx_privacy_engine}-run"
  istanbul_extradata_bootstrap_container_name = "istanbul-extramain-bootstrap"

  consensus_config = {
    raft = {
      geth_args = [
        "--raft",
        "--raftport ${local.raft_port}",
      ]

      enode_params = [
        "raftport=${local.raft_port}",
      ]

      genesis_mixHash    = ["0x00000000000000000000000000000000000000647572616c65787365646c6578"]
      genesis_difficulty = ["0x00"]

      git_url = [""]
    }

    istanbul = {
      geth_args = [
        "--istanbul.blockperiod 1",
        "--emitcheckpoints",
        "--syncmode full",
        "--mine",
        "--minerthreads 1",
      ]

      enode_params = []

      genesis_mixHash    = ["0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365"]
      genesis_difficulty = ["0x01"]

      git_url = ["https://github.com/getamis/istanbul-tools"]
    }
  }

  tx_privacy_address_files = [
    "${var.tx_privacy_engine == "constellation" ? local.constellation_pub_key_file : ""}",
    "${var.tx_privacy_engine == "tessera" ? local.tessera_pub_key_file : ""}",
  ]

  common_container_definitions = [
    "${local.node_key_bootstrap_container_definition}",
    "${local.metadata_bootstrap_container_definition}",
    "${local.quorum_run_container_definition}",
  ]

  container_definitions_for_constellation = [
    "${local.common_container_definitions}",
    "${local.constellation_run_container_definition}",
  ]

  container_definitions_for_tessera = [
    "${local.common_container_definitions}",
    "${local.tessera_run_container_definition}",
  ]

  container_definitions = [
    "${var.tx_privacy_engine == "constellation" ? jsonencode(local.container_definitions_for_constellation) : ""}",
    "${var.tx_privacy_engine == "tessera" ? jsonencode(local.container_definitions_for_tessera) : ""}",
  ]
}

locals {
  constellation_config_file  = "${local.shared_volume_container_path}/constellation.cfg"
  constellation_port         = 10000
  constellation_pub_key_file = "${local.shared_volume_container_path}/tm.pub"

  constellation_config_commands = [
    "constellation-node --generatekeys=${local.shared_volume_container_path}/tm < /dev/null",
    "export HOST_IP=$(cat ${local.host_ip_file})",
    "echo \"\nHost IP: $HOST_IP\"",
    "echo \"Public Key: $(cat ${local.constellation_pub_key_file})\"",
    "all=\"\"; for f in ls ${local.hosts_folder} | grep -v ${local.normalized_host_ip}; do ip=$(cat ${local.hosts_folder}/$f); all=\"$all,\\\"http://$ip:${local.constellation_port}/\\\"\"; done",
    "echo \"Creating ${local.constellation_config_file}\"",
    "echo \"# This file is auto generated. Please do not edit\" > ${local.constellation_config_file}",
    "echo \"url = \\\"http://$HOST_IP:${local.constellation_port}/\\\"\" >> ${local.constellation_config_file}",
    "echo \"port = ${local.constellation_port}\" >> ${local.constellation_config_file}",
    "echo \"socket = \\\"${local.tx_privacy_engine_socket_file}\\\"\" >> ${local.constellation_config_file}",
    "echo \"othernodes = [\\\"http://$HOST_IP:${local.constellation_port}/\\\"$all]\" >> ${local.constellation_config_file}",
    "echo \"publickeys = [\\\"${local.shared_volume_container_path}/tm.pub\\\"]\" >> ${local.constellation_config_file}",
    "echo \"privatekeys = [\\\"${local.shared_volume_container_path}/tm.key\\\"]\" >> ${local.constellation_config_file}",
    "echo \"storage = \\\"/constellation\\\"\" >> ${local.constellation_config_file}",
    "echo \"verbosity = 4\" >> ${local.constellation_config_file}",
    "cat ${local.constellation_config_file}",
  ]

  constellation_run_commands = [
    "set -e",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "${local.constellation_config_commands}",
    "constellation-node ${local.constellation_config_file}",
  ]

  constellation_run_container_definition = {
    name      = "${local.tx_privacy_engine_run_container_name}"
    image     = "${local.tx_privacy_engine_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
      }
    }

    portMappings = [
      {
        hostPort      = "${local.constellation_port}"
        containerPort = "${local.constellation_port}"
      },
    ]

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
    ]

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -S ${local.tx_privacy_engine_socket_file} ];",
      ]
    }

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.constellation_run_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
}

locals {
  quorum_rpc_port                = 22000
  quorum_p2p_port                = 21000
  raft_port                      = 50400
  quorum_data_dir                = "${local.shared_volume_container_path}/dd"
  quorum_password_file           = "${local.shared_volume_container_path}/passwords.txt"
  quorum_static_nodes_file       = "${local.quorum_data_dir}/static-nodes.json"
  quorum_permissioned_nodes_file = "${local.quorum_data_dir}/permissioned-nodes.json"
  genesis_file                   = "${local.shared_volume_container_path}/genesis.json"
  node_id_file                   = "${local.shared_volume_container_path}/node_id"
  node_ids_folder                = "${local.shared_volume_container_path}/nodeids"
  accounts_folder                = "${local.shared_volume_container_path}/accounts"
  privacy_addresses_folder       = "${local.shared_volume_container_path}/privacyaddresses"

  # store Tessera pub keys

  consensus_config_map = "${local.consensus_config[var.consensus_mechanism]}"
  quorum_config_commands = [
    "mkdir -p ${local.quorum_data_dir}/geth",
    "echo \"\" > ${local.quorum_password_file}",
    "echo \"Creating ${local.quorum_static_nodes_file} and ${local.quorum_permissioned_nodes_file}\"",
    "all=\"\"; for f in ls ${local.node_ids_folder}; do nodeid=$(cat ${local.node_ids_folder}/$f); ip=$(cat ${local.hosts_folder}/$f); all=\"$all,\\\"enode://$nodeid@$ip:${local.quorum_p2p_port}?discport=0&${join("&", local.consensus_config_map["enode_params"])}\\\"\"; done; all=$${all:1}",
    "echo \"[$all]\" > ${local.quorum_static_nodes_file}",
    "echo \"[$all]\" > ${local.quorum_permissioned_nodes_file}",
    "echo Permissioned Nodes: $(cat ${local.quorum_permissioned_nodes_file})",
    "geth --datadir ${local.quorum_data_dir} init ${local.genesis_file}",
    "export IDENTITY=$(cat ${local.service_file} | awk -F: '{print $2}')",
  ]
  additional_args = "${local.consensus_config_map["geth_args"]}"
  geth_args = [
    "--datadir ${local.quorum_data_dir}",
    "--rpc",
    "--rpcaddr 0.0.0.0",
    "--rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,${var.consensus_mechanism}",
    "--rpcport ${local.quorum_rpc_port}",
    "--port ${local.quorum_p2p_port}",
    "--unlock 0",
    "--password ${local.quorum_password_file}",
    "--nodiscover",
    "--networkid ${random_integer.network_id.result}",
    "--verbosity 5",
    "--debug",
    "--identity $IDENTITY",
    "--ethstats \"$IDENTITY:${random_id.ethstat_secret.hex}@${aws_instance.bastion.private_ip}:${local.ethstats_port}\"",
  ]
  geth_args_combined = "${join(" ", concat(local.geth_args, local.additional_args))}"
  quorum_run_commands = [
    "set -e",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "echo Wait until ${var.tx_privacy_engine} is ready ...",
    "while [ ! -S \"${local.tx_privacy_engine_socket_file}\" ]; do sleep 1; done",
    "${local.quorum_config_commands}",
    "echo 'Running geth with: ${local.geth_args_combined}'",
    "geth ${local.geth_args_combined}",
  ]
  quorum_run_container_definition = {
    name      = "${local.quorum_run_container_name}"
    image     = "${local.quorum_docker_image}"
    essential = "true"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
      }
    }

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -S ${local.quorum_data_dir}/geth.ipc ];",
      ]
    }

    environments = []

    //portMappings = []
    portMappings = [
      {
        hostPort      = "${local.quorum_rpc_port}"
        containerPort = "${local.quorum_rpc_port}"
      },
      {
        hostPort      = "${local.quorum_p2p_port}"
        containerPort = "${local.quorum_p2p_port}"
      },
      {
        hostPort      = "${local.raft_port}"
        containerPort = "${local.raft_port}"
      },
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
      {
        sourceContainer = "${local.tx_privacy_engine_run_container_name}"
      },
    ]

    environment = [
      {
        name  = "PRIVATE_CONFIG"
        value = "${local.tx_privacy_engine_socket_file}"
      },
      {
        name  = "LD_PRELOAD",
        value = "${local.libfaketime_folder}/libfaketime.so"
      }
    ]

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.quorum_run_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
  genesis = {
    "alloc" = {}

    "coinbase" = "0x0000000000000000000000000000000000000000"

    "config" = {
      "homesteadBlock" = 0
      "byzantiumBlock" = 1
      "chainId"        = "${random_integer.network_id.result}"
      "eip150Block"    = 1
      "eip155Block"    = 0
      "eip150Hash"     = "0x0000000000000000000000000000000000000000000000000000000000000000"
      "eip158Block"    = 1
      "isQuorum"       = "true"
    }

    "difficulty" = "0x0"
    "extraData"  = "0x0000000000000000000000000000000000000000000000000000000000000000"
    "gasLimit"   = "0xE0000000"
    "mixHash"    = "0x00000000000000000000000000000000000000647572616c65787365646c6578"
    "nonce"      = "0x0"
    "parentHash" = "0x0000000000000000000000000000000000000000000000000000000000000000"
    "timestamp"  = "0x00"
  }
}

resource "random_integer" "network_id" {
  min = 2018
  max = 9999

  keepers = {
    changes_when = "${var.network_name}"
  }
}

locals {
  tessera_config_file     = "${local.shared_volume_container_path}/tessera.cfg"
  tessera_port            = 9000
  tessera_thirdparty_port = 9080
  tessera_command         = "java -jar /tessera/tessera-app.jar"
  tessera_pub_key_file    = "${local.shared_volume_container_path}/.pub"

  tessera_config_commands = [
    "apk update",
    "apk add jq",
    "cd ${local.shared_volume_container_path}; echo \"\n\" | ${local.tessera_command} -keygen ${local.shared_volume_container_path}/",
    "export HOST_IP=$(cat ${local.host_ip_file})",
    "export TM_PUB=$(cat ${local.tessera_pub_key_file})",
    "export TM_KEY=$(cat ${local.shared_volume_container_path}/.key)",
    "echo \"\nHost IP: $HOST_IP\"",
    "echo \"Public Key: $TM_PUB\"",
    "all=\"\"; for f in ls ${local.hosts_folder} | grep -v ${local.normalized_host_ip}; do ip=$(cat ${local.hosts_folder}/$f); all=\"$all,{ \\\"url\\\": \\\"http://$ip:${local.tessera_port}/\\\" }\"; done",
    "all=\"[{ \\\"url\\\": \\\"http://$HOST_IP:${local.tessera_port}/\\\" }$all]\"",
    "export TESSERA_VERSION=${var.tessera_docker_image_tag}",
    "export V=$(echo -e \"0.8\n$TESSERA_VERSION\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)",
    "echo \"Creating ${local.tessera_config_file}\"",
    <<SCRIPT
DDIR=${local.quorum_data_dir}
unzip -p /tessera/tessera-app.jar META-INF/MANIFEST.MF | grep Tessera-Version | cut -d: -f2 | xargs
echo "Tessera Version: $TESSERA_VERSION"
V08=$$(echo -e "0.8\n$TESSERA_VERSION" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)
V09=$$(echo -e "0.9\n$TESSERA_VERSION" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)
case "$TESSERA_VERSION" in
    "$V09"|latest)
    # use new config
    cat <<EOF > ${local.tessera_config_file}
{
  "useWhiteList": false,
  "jdbc": {
    "username": "sa",
    "password": "",
    "url": "jdbc:h2:./$${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
    "autoCreateTables": true
  },
  "serverConfigs":[
  {
    "app":"ThirdParty",
    "enabled": true,
    "serverAddress": "http://$HOST_IP:${local.tessera_thirdparty_port}",
    "communicationType" : "REST"
  },
  {
    "app":"Q2T",
    "enabled": true,
    "serverAddress": "unix:${local.tx_privacy_engine_socket_file}",
    "communicationType" : "REST"
  },
  {
    "app":"P2P",
    "enabled": true,
    "serverAddress": "http://$HOST_IP:${local.tessera_port}",
    "sslConfig": {
      "tls": "OFF",
      "generateKeyStoreIfNotExisted": true,
      "serverKeyStore": "$${DDIR}/server-keystore",
      "serverKeyStorePassword": "quorum",
      "serverTrustStore": "$${DDIR}/server-truststore",
      "serverTrustStorePassword": "quorum",
      "serverTrustMode": "TOFU",
      "knownClientsFile": "$${DDIR}/knownClients",
      "clientKeyStore": "$${DDIR}/client-keystore",
      "clientKeyStorePassword": "quorum",
      "clientTrustStore": "$${DDIR}/client-truststore",
      "clientTrustStorePassword": "quorum",
      "clientTrustMode": "TOFU",
      "knownServersFile": "$${DDIR}/knownServers"
    },
    "communicationType" : "REST"
  }
  ],
  "peer": $all,
  "keys": {
    "passwords": [],
    "keyData": [
      {
        "config": $TM_KEY,
        "publicKey": "$TM_PUB"
      }
    ]
  },
  "alwaysSendTo": []
}
EOF    
      ;;
    "$V08")
      # use enhanced config
      cat <<EOF > ${local.tessera_config_file}
{
  "useWhiteList": false,
  "jdbc": {
    "username": "sa",
    "password": "",
    "url": "jdbc:h2:./$${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
    "autoCreateTables": true
  },
  "serverConfigs":[
  {
    "app":"ThirdParty",
    "enabled": true,
    "serverSocket":{
      "type":"INET",
      "port": ${local.tessera_thirdparty_port},
      "hostName": "http://$HOST_IP"
    },
    "communicationType" : "REST"
  },
  {
    "app":"Q2T",
    "enabled": true,
    "serverSocket":{
      "type":"UNIX",
      "path":"${local.tx_privacy_engine_socket_file}"
    },
    "communicationType" : "UNIX_SOCKET"
  },
  {
    "app":"P2P",
    "enabled": true,
    "serverSocket":{
      "type":"INET",
      "port": ${local.tessera_port},
      "hostName": "http://$HOST_IP"
    },
    "sslConfig": {
      "tls": "OFF",
      "generateKeyStoreIfNotExisted": true,
      "serverKeyStore": "$${DDIR}/server-keystore",
      "serverKeyStorePassword": "quorum",
      "serverTrustStore": "$${DDIR}/server-truststore",
      "serverTrustStorePassword": "quorum",
      "serverTrustMode": "TOFU",
      "knownClientsFile": "$${DDIR}/knownClients",
      "clientKeyStore": "$${DDIR}/client-keystore",
      "clientKeyStorePassword": "quorum",
      "clientTrustStore": "$${DDIR}/client-truststore",
      "clientTrustStorePassword": "quorum",
      "clientTrustMode": "TOFU",
      "knownServersFile": "$${DDIR}/knownServers"
    },
    "communicationType" : "REST"
  }
  ],
  "peer": $all,
  "keys": {
    "passwords": [],
    "keyData": [
      {
        "config": $TM_KEY,
        "publicKey": "$TM_PUB"
      }
    ]
  },
  "alwaysSendTo": []
}
EOF
      ;;
    *)
    # use old config
    cat <<EOF > ${local.tessera_config_file}
{
    "useWhiteList": false,
    "jdbc": {
        "username": "sa",
        "password": "",
        "url": "jdbc:h2:./$${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
        "autoCreateTables": true
    },
    "server": {
        "port": 9000,
        "hostName": "http://$HOST_IP",
        "sslConfig": {
            "tls": "OFF",
            "generateKeyStoreIfNotExisted": true,
            "serverKeyStore": "$${DDIR}/server-keystore",
            "serverKeyStorePassword": "quorum",
            "serverTrustStore": "$${DDIR}/server-truststore",
            "serverTrustStorePassword": "quorum",
            "serverTrustMode": "TOFU",
            "knownClientsFile": "$${DDIR}/knownClients",
            "clientKeyStore": "$${DDIR}/client-keystore",
            "clientKeyStorePassword": "quorum",
            "clientTrustStore": "$${DDIR}/client-truststore",
            "clientTrustStorePassword": "quorum",
            "clientTrustMode": "TOFU",
            "knownServersFile": "$${DDIR}/knownServers"
        }
    },
    "peer": $all,
    "keys": {
        "passwords": [],
        "keyData": [
            {
                "config": $TM_KEY,
                "publicKey": "$TM_PUB"
            }
        ]
    },
    "alwaysSendTo": [],
    "unixSocketFile": "${local.tx_privacy_engine_socket_file}"
}
EOF
      ;;
esac
cat ${local.tessera_config_file}
SCRIPT
    ,
  ]

  tessera_run_commands = [
    "set -e",
    "echo Wait until metadata bootstrap completed ...",
    "while [ ! -f \"${local.metadata_bootstrap_container_status_file}\" ]; do sleep 1; done",
    "${local.tessera_config_commands}",
    "${local.tessera_command} -configfile ${local.tessera_config_file}",
  ]

  tessera_run_container_definition = {
    name      = "${local.tx_privacy_engine_run_container_name}"
    image     = "${local.tx_privacy_engine_docker_image}"
    essential = "false"

    logConfiguration = {
      logDriver = "awslogs"

      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.quorum.name}"
        awslogs-region        = "${var.region}"
        awslogs-stream-prefix = "logs"
      }
    }

    portMappings = [
      {
        hostPort      = "${local.tessera_port}"
        containerPort = "${local.tessera_port}"
      },
      {
        hostPort      = "${local.tessera_thirdparty_port}"
        containerPort = "${local.tessera_thirdparty_port}"
      },
    ]

    mountPoints = [
      {
        sourceVolume  = "${local.shared_volume_name}"
        containerPath = "${local.shared_volume_container_path}"
      },
    ]

    volumesFrom = [
      {
        sourceContainer = "${local.metadata_bootstrap_container_name}"
      },
    ]

    healthCheck = {
      interval    = 30
      retries     = 10
      timeout     = 60
      startPeriod = 300

      command = [
        "CMD-SHELL",
        "[ -S ${local.tx_privacy_engine_socket_file} ];",
      ]
    }

    entrypoint = [
      "/bin/sh",
      "-c",
      "${join("\n", local.tessera_run_commands)}",
    ]

    dockerLabels = "${local.common_tags}"

    cpu = 0
  }
}

locals {
  service_name_fmt = "node-%0${min(length(format("%d", var.number_of_nodes)), length(format("%s", var.number_of_nodes))) + 1}d-%s"
  ecs_cluster_name = "quorum-network-${var.network_name}"
  quorum_bucket    = "${var.region}-ecs-${lower(var.network_name)}-${random_id.bucket_postfix.hex}"
}

resource "aws_ecs_cluster" "quorum" {
  name = "${local.ecs_cluster_name}"
}

resource "aws_ecs_task_definition" "quorum" {
  family                   = "quorum-${var.consensus_mechanism}-${var.tx_privacy_engine}-${var.network_name}"
  container_definitions    = "${replace(element(compact(local.container_definitions), 0), "/\"(true|false|[0-9]+)\"/", "$1")}"
  requires_compatibilities = ["${var.ecs_mode}"]
  cpu                      = "4096"
  memory                   = "8192"
  network_mode             = "${var.ecs_network_mode}"
  task_role_arn            = "${aws_iam_role.ecs_task.arn}"
  execution_role_arn       = "${aws_iam_role.ecs_task.arn}"

  volume {
    name = "${local.shared_volume_name}"
  }
}

resource "aws_ecs_service" "quorum" {
  count           = "${var.number_of_nodes}"
  name            = "${format(local.service_name_fmt, count.index + 1, var.network_name)}"
  cluster         = "${aws_ecs_cluster.quorum.id}"
  task_definition = "${aws_ecs_task_definition.quorum.arn}"
  launch_type     = "EC2"
  desired_count   = "1"

  // not compatible with 'bridge' network mode
  //network_configuration {
  //  subnets          = ["${var.subnet_ids}"]
  //  assign_public_ip = "${var.is_igw_subnets}"
  //  security_groups  = ["${aws_security_group.quorum.id}"]
  //}
}

data "aws_caller_identity" "this" {}

data "aws_iam_policy_document" "kms_policy" {
  statement {
    sid = "AllowAccess"

    actions = [
      "kms:*",
    ]

    effect = "Allow"

    resources = ["*"]

    principals {
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.this.account_id}:root",
      ]

      type = "AWS"
    }
  }
}

resource "aws_kms_key" "bucket" {
  description             = "Used to encrypt/decrypt objects stored inside bucket created for this deployment"
  policy                  = "${data.aws_iam_policy_document.kms_policy.json}"
  deletion_window_in_days = "7"
  tags                    = "${local.common_tags}"
}

resource "random_id" "bucket_postfix" {
  byte_length = 8
}

data "aws_iam_policy_document" "bucket_policy" {
  statement {
    sid     = "AllowAccess"
    actions = ["s3:*"]
    effect  = "Allow"

    resources = [
      "arn:aws:s3:::${local.quorum_bucket}",
      "arn:aws:s3:::${local.quorum_bucket}/*",
    ]

    principals {
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.this.account_id}:root"]
      type        = "AWS"
    }
  }

  statement {
    sid     = "DenyAccess1"
    actions = ["s3:PutObject"]
    effect  = "Deny"

    resources = [
      "arn:aws:s3:::${local.quorum_bucket}",
      "arn:aws:s3:::${local.quorum_bucket}/*",
    ]

    principals {
      identifiers = ["*"]
      type        = "AWS"
    }

    condition {
      test     = "Null"
      values   = ["true"]
      variable = "s3:x-amz-server-side-encryption"
    }
  }

  statement {
    sid     = "DenyAccess2"
    actions = ["s3:PutObject"]
    effect  = "Deny"

    resources = [
      "arn:aws:s3:::${local.quorum_bucket}",
      "arn:aws:s3:::${local.quorum_bucket}/*",
    ]

    principals {
      identifiers = ["*"]
      type        = "AWS"
    }

    condition {
      test     = "StringNotEquals"
      values   = ["aws:kms"]
      variable = "s3:x-amz-server-side-encryption"
    }
  }
}

resource "aws_s3_bucket" "quorum" {
  bucket        = "${local.quorum_bucket}"
  region        = "${var.region}"
  policy        = "${data.aws_iam_policy_document.bucket_policy.json}"
  force_destroy = true

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    "rule" {
      "apply_server_side_encryption_by_default" {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = "${aws_kms_key.bucket.arn}"
      }
    }
  }
}

resource "aws_iam_role" "ecs_task" {
  name = "quorum-ecs-task-${var.network_name}"
  path = "/ecs/"

  assume_role_policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": [
          "ecs.amazonaws.com",
          "ecs-tasks.amazonaws.com"
        ]
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "ecs_task" {
  statement {
    sid = "AllowS3Access"

    actions = [
      "s3:*",
    ]

    resources = [
      "arn:aws:s3:::${local.quorum_bucket}",
      "arn:aws:s3:::${local.quorum_bucket}/*",
    ]
  }

  statement {
    sid = "AllowKMSAccess"

    actions = [
      "kms:*",
    ]

    resources = [
      "${aws_kms_key.bucket.arn}",
    ]
  }

  statement {
    sid = "AllowECS"

    actions = [
      "ecs:DescribeTasks",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    sid = "AllowS3Bastion"

    actions = [
      "s3:*",
    ]

    resources = [
      "arn:aws:s3:::${local.bastion_bucket}",
      "arn:aws:s3:::${local.bastion_bucket}/*",
    ]
  }
}

resource "aws_iam_policy" "ecs_task" {
  name        = "quorum-ecs-task-policy-${var.network_name}"
  path        = "/"
  description = "This policy allows task to access S3 bucket"
  policy      = "${data.aws_iam_policy_document.ecs_task.json}"
}

resource "aws_iam_role_policy_attachment" "ecs_task_s3" {
  role       = "${aws_iam_role.ecs_task.id}"
  policy_arn = "${aws_iam_policy.ecs_task.arn}"
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution" {
  role       = "${aws_iam_role.ecs_task.id}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "ecs_task_cloudwatch" {
  role       = "${aws_iam_role.ecs_task.id}"
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

resource "aws_cloudwatch_log_group" "quorum" {
  name              = "/ecs/quorum/${var.network_name}"
  retention_in_days = "7"
  tags              = "${local.common_tags}"
}

provider "null" {
  version = "~> 1.0"
}

provider "random" {
  version = "~> 2.0"
}

provider "local" {
  version = "~> 1.1"
}

provider "tls" {
  version = "~> 1.2"
}

locals {
  tessera_docker_image           = "${var.tx_privacy_engine == "tessera" ? format("%s:%s", var.tessera_docker_image, var.tessera_docker_image_tag) : ""}"
  constellation_docker_image     = "${var.tx_privacy_engine == "constellation" ? format("%s:%s", var.constellation_docker_image, var.constellation_docker_image_tag) : ""}"
  quorum_docker_image            = "${format("%s:%s", var.quorum_docker_image, var.quorum_docker_image_tag)}"
  tx_privacy_engine_docker_image = "${coalesce(local.tessera_docker_image, local.constellation_docker_image)}"
  aws_cli_docker_image           = "${format("%s:%s", var.aws_cli_docker_image, var.aws_cli_docker_image_tag)}"

  common_tags = {
    "NetworkName"               = "${var.network_name}"
    "ECSClusterName"            = "${local.ecs_cluster_name}"
    "DockerImage.Quorum"        = "${local.quorum_docker_image}"
    "DockerImage.PrivacyEngine" = "${local.tx_privacy_engine_docker_image}"
  }
}

output "_status" {
  value = <<MSG
Completed!

Quorum Docker Image         = ${local.quorum_docker_image}
Privacy Engine Docker Image = ${local.tx_privacy_engine_docker_image}
Number of Quorum Nodes      = ${var.number_of_nodes}
ECS Task Revision           = ${aws_ecs_task_definition.quorum.revision}
CloudWatch Log Group        = ${aws_cloudwatch_log_group.quorum.name}
MSG
}

output "bastion_host_dns" {
  value = "${aws_instance.bastion.public_dns}"
}

output "bastion_host_ip" {
  value = "${aws_instance.bastion.public_ip}"
}

output "network_name" {
  value = "${var.network_name}"
}

output "ecs_cluster_name" {
  value = "${aws_ecs_cluster.quorum.name}"
}

output "chain_id" {
  value = "${random_integer.network_id.result}"
}

output "private_key_file" {
  value = "${local_file.private_key.filename}"
}

output "bucket_name" {
  value = "${aws_s3_bucket.quorum.bucket}"
}

resource "aws_security_group" "quorum" {
  vpc_id      = "${local.vpc_id}"
  name        = "quorum-sg-${var.network_name}"
  description = "Security group used in Quorum network ${var.network_name}"

  egress {
    from_port = 0
    protocol  = "-1"
    to_port   = 0

    cidr_blocks = [
      "0.0.0.0/0",
    ]

    description = "Allow all"
  }

  tags = "${merge(local.common_tags, map("Name", format("quorum-sg-%s", var.network_name)))}"
}

resource "aws_security_group_rule" "ethstats" {
  from_port         = "${local.ethstats_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.ethstats_port}"
  type              = "ingress"
  self              = true
  description       = "ethstats traffic"
}

resource "aws_security_group_rule" "geth_p2p" {
  from_port         = "${local.quorum_p2p_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.quorum_p2p_port}"
  type              = "ingress"
  self              = true
  description       = "Geth P2P traffic"
}

resource "aws_security_group_rule" "geth_admin_rpc" {
  from_port         = "${local.quorum_rpc_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.quorum_rpc_port}"
  type              = "ingress"
  self              = "true"
  description       = "Geth Admin RPC traffic"
}

resource "aws_security_group_rule" "constellation" {
  count             = "${var.tx_privacy_engine == "constellation" ? 1 : 0}"
  from_port         = "${local.constellation_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.constellation_port}"
  type              = "ingress"
  self              = "true"
  description       = "Constellation API traffic"
}

resource "aws_security_group_rule" "tessera" {
  count             = "${var.tx_privacy_engine == "tessera" ? 1 : 0}"
  from_port         = "${local.tessera_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.tessera_port}"
  type              = "ingress"
  self              = "true"
  description       = "Tessera API traffic"
}

resource "aws_security_group_rule" "tessera_thirdparty" {
  count             = "${var.tx_privacy_engine == "tessera" ? 1 : 0}"
  from_port         = "${local.tessera_thirdparty_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.tessera_thirdparty_port}"
  type              = "ingress"
  self              = "true"
  description       = "Tessera Thirdparty API traffic"
}

resource "aws_security_group_rule" "raft" {
  count             = "${var.consensus_mechanism == "raft" ? 1 : 0}"
  from_port         = "${local.raft_port}"
  protocol          = "tcp"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = "${local.raft_port}"
  type              = "ingress"
  self              = "true"
  description       = "Raft HTTP traffic"
}

resource "aws_security_group_rule" "open-all-ingress-research" {
  count             = "${var.ecs_mode == "EC2" ? 1 : 0}"
  from_port         = 0
  protocol          = "-1"
  security_group_id = "${aws_security_group.quorum.id}"
  to_port           = 0
  type              = "ingress"
  cidr_blocks       = ["${var.access_ec2_nodes_cidr_blocks}"]
  description       = "Open all ports"
}

variable "region" {
  description = "Target AWS Region. This must be pre-initialized from _terraform_init run"
  default = "us-east-1"
}

variable "network_name" {
  description = "Identify the Quorum network from multiple deployments. This must be pre-initialized from _terraform_init run"
  default = "apollo"
}

variable "number_of_nodes" {
  description = "Number of Quorum nodes. Default is 7"
  default     = "0"
}

variable "asg_instance_type" {
  description = "ASG instance type for EC2 based quorum"
  default     = "t2.xlarge"
}

variable "ecs_mode" {
  description = "ECS engine mode: EC2 or FARGATE"
  default     = "EC2"
}

variable "ecs_network_mode" {
  description = "ECS network node: awsvpc or bridge"
  default     = "bridge"
}

variable "client_name" {
  description = "Etherum client name"
  default     = "quorum"
}

variable "quorum_docker_image" {
  description = "URL to Quorum docker image to be used"
  default     = "quorumengineering/quorum"
}

variable "quorum_docker_image_tag" {
  description = "Quorum Docker image tag to be used"
  default     = "latest"
}

variable "constellation_docker_image" {
  description = "URL to Constellation docker image to be used. Only needed if tx_privacy_engine is constellation"
  default     = "quorumengineering/constellation"
}

variable "constellation_docker_image_tag" {
  description = "Constellation Docker image tag to be used"
  default     = "latest"
}

variable "tessera_docker_image" {
  description = "URL to Constellation docker image to be used. Only needed if tx_privacy_engine is constellation"
  default     = "quorumengineering/tessera"
}

variable "tessera_docker_image_tag" {
  description = "Tessera Docker image tag to be used"
  default     = "latest"
}

variable "aws_cli_docker_image" {
  description = "To interact with AWS services"
  default     = "senseyeio/alpine-aws-cli"
}

variable "aws_cli_docker_image_tag" {
  description = "AWS CLI Docker image tag to be used"
  default     = "latest"
}

variable "consensus_mechanism" {
  description = "Concensus mechanism used in the network. Supported values are raft/istanbul"
  default     = "istanbul"
}

variable "tx_privacy_engine" {
  description = "Engine that implements transaction privacy. Supported values are constellation/tessera"
  default     = "tessera"
}

//TODO remove
/*
variable "quorum_bucket" {
  description = "This is to store shared data during the bootstrap. This must be pre-initialized from _terraform_init run"
}

variable "quorum_bucket_kms_key_arn" {
  description = "To encrypt/decrypt objects stored in quorum_bucket. This must be pre-initialized from _terraform_init run"
}
*/
variable "access_bastion_cidr_blocks" {
  type        = "list"
  description = "CIDR blocks that will be added to allow SSH to Bastion Node"
  default     = ["0.0.0.0/0"]
}

variable "access_ec2_nodes_cidr_blocks" {
  type        = "list"
  description = "CIDR blocks that will be added to allow all traffic to cluster EC2 nodes"
  default     = ["0.0.0.0/0"]
}

//variable bucket {}
variable profile {
  default = "default"
}

variable "faketime" {
  type    = "list"
  description = "A faketime value passed to cluster node"
  default = [ "1", "-3", "2" ]

}
`
)
