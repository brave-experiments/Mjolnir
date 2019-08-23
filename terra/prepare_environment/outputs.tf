output "arn" {
  value = "${aws_kms_key.s3.arn}"
}

output "id" {
  value = "${aws_kms_key.s3.key_id}"
}