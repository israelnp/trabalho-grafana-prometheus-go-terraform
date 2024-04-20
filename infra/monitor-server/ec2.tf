resource "aws_instance" "this" {
  ami                    = var.image_id 
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.this.id]
  key_name               = var.key_name
  user_data = base64encode(
  templatefile("monitor-setup.sh", {
    target-ip = var.target-ip
  }))
  monitoring             = true
  subnet_id              = var.aws_subnet_id
  associate_public_ip_address = false
  tags = merge(var.common_tags, { Name = "Monitor Machine" })
}