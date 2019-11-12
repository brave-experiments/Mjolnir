variable "region" {
  description = "Target AWS Region"
  default = "us-east-1"
}

variable "network_name" {
  description = "Identify the Quorum network from multiple deployments"
  default = "mjolnir"
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

variable "pantheon_docker_image" {
  description = "URL to Pantheon docker image to be used"
  default     = "pegasyseng/pantheon"
}

variable "pantheon_docker_image_tag" {
  description = "Pantheon Docker image tag to be used"
  default     = "latest"
}

variable "parity_docker_image" {
  description = "URL to Parity HBBFT docker image to be used"
  default     = "brave/honey-badger"
}


variable "parity_docker_image_tag" {
  description = "Parity Docker image tag to be used"
  default     = "latest"
}

variable "honey_badger_config_gen" {
  description = "URL to Parity HBBFT docker image to be used"
  default     = "brave/honey-badger-config-generator"
}

variable "honey_badger_config_gen_tag" {
  description = "Honey Badger Docker image tag to be used"
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

variable "chaos_testing_docker_image" {
  description = "Chaos testing tool docker image to be used"
  default     = "gaiaadm/pumba"
}

variable "chaos_testing_docker_image_tag" {
  description = "Chaos testing tool Docker image tag to be used"
  default     = "latest"
}

variable "chaos_testing_run_command" {
  description = "Chaos testing tool run command to be used"
  default     = []
}

variable "consensus_mechanism" {
  description = "Concensus mechanism used in the network. Supported values are raft/istanbul"
  default     = "istanbul"
}

variable "tx_privacy_engine" {
  description = "Engine that implements transaction privacy. Supported values are constellation/tessera"
  default     = "tessera"
}

variable "genesis_gas_limit" {
  description = "Gas limit parameter across all clients"
  default = "0xE0000000"
}

variable "genesis_timestamp" {
  description = "Epoch timestamp used for genesis file"
  default = "0x00"
}

variable "genesis_difficulty" {
  description = "Difficulty used for genesis file"
  default = "0x0"
}

variable "genesis_nonce" {
  description = "Nonce used for genesis file"
  default = "0x0"
}

variable "genesis_blocktime" {
  description = "Blocktime parameter across all clients"
  default = "5"
}

variable "genesis_min_gas_limit" {
  description = "minGasLimit used for parity genesis file"
  default = "0x1388"
}

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
  default = []

}

variable "chainhammer_repo_url" {
   description = "TPS measurements of parity aura, geth clique, quorum, tobalaba, etc"
   default = "https://github.com/drandreaskrueger/chainhammer.git"
}

variable "bastion_volume_size" {
  description = "Bastion root volume size in GB"
  default = 32
}
