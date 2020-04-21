variable "ssh_public_key_file" {
  type = string
  description = "Path to public SSH key file to use for AWS Instances"
}

variable "staging_instance_type" {
  type = string
  description = "Instance type for staging"
  default = "t2.micro"
}
