terraform {
  required_providers {
    instantproxy = {
      version = "0.1"
      source  = "local/dthung1602/instantproxy"
    }
  }
}

provider "instantproxy" {
  endpoint = "http://localhost:3000"  // run
  username = "username"
  password = "password"
}

data "instantproxy_proxies" "proxies" {}

output "proxies" {
  value = data.instantproxy_proxies.proxies
}

resource "instantproxy_authorized_ips" "auth_ips" {
  value = ["127.0.0.1"]
}

output "auth_ips" {
  value = resource.instantproxy_authorized_ips.auth_ips
}
