provider "aws" {
  version = "~> 2.0"
  region  = "eu-west-2"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_key_pair" "app-key" {
  key_name = "app"
  public_key = file(var.ssh_public_key_file)
}

module "staging" {
  source = "./application"

  instance_type = "t2.micro"
  instance_ami = data.aws_ami.ubuntu.id
  instance_count = 1
  instance_key_name = aws_key_pair.app-key.key_name
  stage = "staging"
}

module  "production" {
  source = "./application"

  instance_type = "t2.micro"
  instance_ami = data.aws_ami.ubuntu.id
  instance_count = 2
  instance_key_name = aws_key_pair.app-key.key_name
  stage = "production"
}
