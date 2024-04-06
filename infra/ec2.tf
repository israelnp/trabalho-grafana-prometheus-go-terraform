#Configurações da instancia
resource "aws_instance" "this" {
  ami                         = var.instance_ami
  instance_type               = var.instance_type
  associate_public_ip_address = true
  key_name                    = var.instance_key_name

  user_data = base64encode(
    templatefile("setup.sh",
      {
        ///pode passar alguma valor como variável de ambiente aqui

  }))
  monitoring             = true
  subnet_id              = aws_subnet.this["pub_a"].id
  tags                   = { Name = "APP GO" }
  vpc_security_group_ids = [aws_security_group.go.id]


}