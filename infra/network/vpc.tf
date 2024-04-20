resource "aws_vpc" "this" {
  cidr_block = var.vpc_cidr_block
  enable_dns_hostnames = true
  tags       = merge(var.common_tags, { Name = "Terraform VPC " })
}


output vpc_id {
  value       = aws_vpc.this.id
}