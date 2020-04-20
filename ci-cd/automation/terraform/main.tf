provider "aws" {
  version     = "~> 2.49"
  region      = "eu-west-3"
}

resource "aws_key_pair" "ci-cd-key" {
  key_name      = "ci-cd"
  public_key    = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvt61yWq1S0PfDRz/oK77YyT7q1XKJ1WbLPKA5NN7UiNh50zpPhqlr1fDnOHv2mn5u2k8X2aRnlvZlP+vHY2ExrrLN/gaeLFUP7eEw/KAw8+Y50WGDHXD6fCuHDI2VvwQ/gsZtMkiXD1Faj2E4CMfOfuXt+CORqLFxd1kMC4f9dgB+mBW+ZtqEtJt6Urql23lZeoE6aSJBSc9Ev29mX8QafAAVtWX72ca1IEZ8MVykIrFDGtOjAQ2tAKSaQkX4B0OBMpxvaybe0kHQlK++oFs4awWyRrlXPppbK2qWXy9bGwe7bxuJtP2FO7ot4WM4TdqnpvBBXpuFDQm8Gtt3+zdb antoinemasselot@Antoines-MacBook-Pro-2.local"
}

resource "aws_instance" "ci-cd" {
  key_name      = "${aws_key_pair.ci-cd-key.key_name}"
  ami           = "ami-096b8af6e7e8fb927"
  instance_type = "t2.micro"

  security_groups = ["${aws_security_group.web_group_security.name}"]
}

resource "aws_security_group" "web_group_security" {
  name        = "madu_security_group"
  description = "security group for the project Madu"

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 63000
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


