variable "aws_region" {
  type        = string
  description = ""
  default     = "us-east-1"
}

variable "vpc_cidr_block" {
  type        = string
  description = ""
  default     = "192.168.0.0/16"
}

variable "az_subnets_cidr_block" {
  type        = map(any)
  description = ""
  default = {
    "pub_a" : ["192.168.1.0/24", "us-east-1a", "Public A"]
    "pub_b" : ["192.168.2.0/24", "us-east-1b", "Public B"]
    "pvt_a" : ["192.168.3.0/24", "us-east-1a", "Private A"]
    "pvt_b" : ["192.168.4.0/24", "us-east-1b", "Private B"]
  }
}

variable "aws_profile" {
  type        = string
  description = ""
  default     = "default"
}
variable "service_name" {
  type        = string
  description = ""
  default     = "autoscaling-wordpress"
}

variable "instance_type" {
  type        = string
  description = ""
  default     = "t3.large"
}

variable "instance_key_name" {
  type        = string
  description = ""
  default     = "grafanademo"
}

variable "name_prefix" {
  type        = string
  description = ""
  default     = "terraform-"
}

variable "image_id" {
  type        = string
  description = ""
  default     = "ami-04505e74c0741db8d"
}
