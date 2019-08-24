provider "aws" {
  region  = "${var.region}"
  version = "~> 1.36"
  profile = "${var.profile}"
}

resource "random_id" "kms_key_alias" {
  byte_length = 8
}

data "aws_caller_identity" "current" {}

resource "aws_s3_bucket" "terraform_remote_state" {
  bucket        = "${var.region}-${var.client_name}-deployment-state-${var.network_name}"
  acl           = "private"
  force_destroy = true

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = "${aws_kms_key.s3.arn}"
        sse_algorithm     = "aws:kms"
      }
    }
  }

  tags = {
    Name        = "terraform-remote-state"
    Environment = "${var.network_name}"
  }
}

resource "aws_s3_bucket_policy" "terraform_remote_state" {
  bucket = "${aws_s3_bucket.terraform_remote_state.id}"
  policy =<<EOF
{
  "Version": "2012-10-17",
  "Id": "RequireEncryption",
   "Statement": [
    {
      "Sid": "Allow access for IAM Users",
      "Effect": "Allow",
      "Action": ["s3:*"],
      "Resource": ["arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}","arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}/*"],
      "Principal": {
        "AWS": "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      }
    },
    {
      "Sid": "Deny access when no encryption header",
      "Effect": "Deny",
      "Action": ["s3:PutObject"],
      "Resource": ["arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}","arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}/*"],
      "Condition": {
        "Null": {
          "s3:x-amz-server-side-encryption": "true"
        }
      },
      "Principal": "*"
    },
    {
      "Sid": "Deny access when no aws:kms encryption algo header",
      "Effect": "Deny",
      "Action": ["s3:PutObject"],
      "Resource": ["arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}","arn:aws:s3:::${aws_s3_bucket.terraform_remote_state.bucket}/*"],
      "Condition": {
        "StringNotEquals": {
          "s3:x-amz-server-side-encryption": "aws:kms"
        }
      },
      "Principal": "*"
    }
  ]
}
EOF
}

resource "aws_kms_key" "s3" {
  description             = "KMS key for s3"
  deletion_window_in_days = 7
  enable_key_rotation     = true

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Id": "key-default-1",
  "Statement": [
    {
      "Sid": "Enable IAM User Permissions",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      },
      "Action": "kms:*",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_kms_alias" "s3" {
  name          = "alias/${var.region}-${var.client_name}-${random_id.kms_key_alias.dec}-deployment-state"
  target_key_id = "${aws_kms_key.s3.key_id}"
}

//state_init

locals {
  tfinit_filename = "terraform.auto.backend_config"
  tfvars_filename = "terraform.auto.tfvars"
  deployment_id   = "${coalesce(var.network_name, join("", random_id.deployment_id.*.b64_url))}"
}

resource "random_id" "deployment_id" {
  count       = 1
  prefix      = "q-"
  byte_length = 8
}

resource "local_file" "tfinit" {
  filename = "${format("%s/../shared/%s", path.module, local.tfinit_filename)}"

  content = <<EOF
# This file is auto generated. Please do not edit
# This is the backend configuration for `terraform init` in the main deployment
region="${var.region}"
bucket="${aws_s3_bucket.terraform_remote_state.id}"
encrypt="true"
kms_key_id="${aws_kms_key.s3.arn}"
key="${local.deployment_id}"
profile="${var.profile}"
EOF
}

resource "local_file" "tfvars" {
  filename = "${format("%s/../shared/%s", path.module, local.tfvars_filename)}"

  content = <<EOF
# This file is auto generated. Please do not edit
# This file contains the default values for required variables
region="${var.region}"
network_name="${local.deployment_id}"
${var.client_name}_bucket="${aws_s3_bucket.terraform_remote_state.id}"
${var.client_name}_bucket_kms_key_arn="${aws_kms_key.s3.arn}"
profile="${var.profile}"
EOF
}
