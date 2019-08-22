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
