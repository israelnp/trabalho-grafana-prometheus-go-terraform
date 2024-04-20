resource "aws_subnet" "this" {
  for_each = var.az_subnets_cidr_block

  vpc_id            = aws_vpc.this.id
  cidr_block        = each.value[0]
  availability_zone = each.value[1]

  tags = merge(var.common_tags, { Name = each.value[2] })
}


output  public_az_a_subnet_id {
  value       = aws_subnet.this["pub_a"].id
}

output  public_az_b_subnet_id {
  value       = aws_subnet.this["pub_b"].id
}

output  private_az_a_subnet_id {
  value       = aws_subnet.this["pvt_a"].id
}

output  private_az_b_subnet_id {
  value       = aws_subnet.this["pvt_b"].id
}