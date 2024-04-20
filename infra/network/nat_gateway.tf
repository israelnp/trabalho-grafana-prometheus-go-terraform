resource "aws_eip" "this" {
  vpc = true
}
resource "aws_nat_gateway" "this" {
  allocation_id = aws_eip.this.id
  subnet_id = aws_subnet.this["pub_a"].id
  tags   = merge(var.common_tags, { Name = "Terraform NGW " })
}