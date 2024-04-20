variable az_subnets_cidr_block {
  type        = map
  description = "Wordpress AZ subnets cidr_block"
}

variable vpc_cidr_block {
  type        = string
  description = "Wordpress vpc cidr_block"
}

variable common_tags {
  type        = map
  description = ""
}
