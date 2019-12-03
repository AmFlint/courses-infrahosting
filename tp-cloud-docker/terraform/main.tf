
provider "scaleway" {
  access_key = ""
  secret_key = ""
  organization_id = ""
  zone       = "fr-par-1"
  region     = "fr-par"
}

# Adresse IP Publique à lier à l'instance

resource "scaleway_instance_ip" "public_ip" {
  server_id = scaleway_instance_server.web.id
}

# Security group

resource "scaleway_instance_security_group" "www" {
  inbound_default_policy = "drop"
  outbound_default_policy = "accept"

  inbound_rule {
    action = "accept"
    port = "22"
    ip = "78.31.41.54"
  }

  inbound_rule {
    action = "accept"
    port = "80"
  }

  inbound_rule {
    action = "accept"
    port = "8080"
    ip = "78.31.41.54"
  }

  inbound_rule {
    action = "accept"
    port = "3000"
  }
}

# Instance

resource "scaleway_instance_server" "web" {
  type = "DEV1-S"
  image = "f974feac-abae-4365-b988-8ec7d1cec10d"

  tags = [ "application" ]

  security_group_id= scaleway_instance_security_group.www.id
}
