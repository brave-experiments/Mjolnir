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
      "arn:aws:s3:::${local.parity_bucket}",
      "arn:aws:s3:::${local.parity_bucket}/*",
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
  name        = "parity-bastion-policy-${var.network_name}"
  path        = "/"
  description = "This policy allows task to access S3 bucket and ECS"
  policy      = "${data.aws_iam_policy_document.bastion.json}"
}

resource "aws_iam_role_policy_attachment" "bastion" {
  role       = "${aws_iam_role.bastion.id}"
  policy_arn = "${aws_iam_policy.bastion.arn}"
}
