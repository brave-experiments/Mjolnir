provider "aws" {
  region  = "${var.region}"
  version = "~> 1.36"
  profile = "${var.profile}"
}

data "aws_caller_identity" "current" {}

resource "aws_s3_bucket" "terraform_remote_state" {
  bucket = "${var.region}-quorum-deployment-state-${var.network_name}"
  acl    = "private"

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
  name          = "alias/${var.region}-quorum-deployment-state"
  target_key_id = "${aws_kms_key.s3.key_id}"
}

