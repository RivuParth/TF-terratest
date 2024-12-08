module "s3" {
  source      = "./modules/s3"
  bucket_name = var.bucket_name
  bucket_acl  = var.bucket_acl
}

output "bucket_id" {
  value = module.s3.bucket_id
}
