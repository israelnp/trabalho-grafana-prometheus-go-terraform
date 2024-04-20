resource "aws_instance" "this" {
  ami                    = "ami-080e1f13689e07408"
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.this.id]
  key_name               = var.key_name
  user_data = base64encode(
  templatefile("vpn-setup.sh", {}))
  monitoring                  = true
  subnet_id                   = var.aws_subnet_id
  associate_public_ip_address = true
  tags                        = merge(var.common_tags, { Name = "VPN Machine" })
}

