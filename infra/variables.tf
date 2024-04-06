variable "aws_region" {
  type        = string
  description = ""
  default     = "us-east-1"
}

variable "aws_profile" {
  type        = string
  description = ""
  default     = "default"
}

variable "instance_type" {
  type        = string
  description = ""
  default     = "t3.medium"
}

variable "instance_ami" {
  type        = string
  description = ""
  default     = "ami-04505e74c0741db8d"
}

variable "instance_key_name" {
  type        = string
  description = ""
  default     = "grafanademo"
}


