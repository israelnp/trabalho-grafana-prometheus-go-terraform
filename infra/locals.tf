locals {
  subnet_ids = { for k, v in aws_subnet.this : v.tags.Name => v.id }

  common_tags = {
    Project   = "Trabalho Fundamentos do DevOps"
    CreatedAt = "2024-14-06"
    ManagedBy = "Terraform"
    Service   = "APP GO"
  }
}
