resource "aws_instance" "this" {
  ami                    = var.image_id
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.this.id]
  key_name               = var.key_name
  user_data = base64encode(
  templatefile("app-setup.sh", {}))
  monitoring                  = true
  subnet_id                   = var.aws_subnet_id
  associate_public_ip_address = true
  tags                        = merge(var.common_tags, { Name = "APP Machine" })
}


output app-ip {
  value       = aws_instance.this.private_ip
}
