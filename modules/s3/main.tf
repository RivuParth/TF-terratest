resource "aws_s3_bucket" "test_bucket" {
  bucket = var.bucket_name
  acl    = var.bucket_acl
}

output "bucket_id" {
  value = aws_s3_bucket.test_bucket.id
}
