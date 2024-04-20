terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.32.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }
}

provider "aws" {
  access_key = var.access_key
  secret_key = var.secret_key
  region     = var.aws_region
  profile    = var.aws_profile
}

module "network" {
  source                = "./network"
  vpc_cidr_block        = var.vpc_cidr_block
  az_subnets_cidr_block = var.az_subnets_cidr_block
  common_tags           = local.common_tags
}

module "vpnServer" {
  source        = "./vpn-server"
  image_id      = var.image_id
  instance_type = var.instance_type
  aws_subnet_id = module.network.public_az_a_subnet_id
  key_name      = var.instance_key_name
  vpc_id        = module.network.vpc_id
  common_tags   = local.common_tags
}

module "appServer" {
  source        = "./app-server"
  image_id      = var.image_id
  instance_type = var.instance_type
  aws_subnet_id = module.network.public_az_a_subnet_id
  key_name      = var.instance_key_name
  vpc_id        = module.network.vpc_id
  common_tags   = local.common_tags
}

module "monitorServer" {
  source        = "./monitor-server"
  image_id      = var.image_id
  instance_type = var.instance_type
  aws_subnet_id = module.network.private_az_a_subnet_id
  key_name      = var.instance_key_name
  vpc_id        = module.network.vpc_id
  common_tags   = local.common_tags
  target-ip     = module.appServer.app-ip
}



