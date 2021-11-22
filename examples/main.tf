terraform {
  required_providers {
    gocache = {
      source = "gocache.com.br/gocachebr/gocache"
      version = "0.1.0"
    }
  }
}

provider "gocache" {
  # token pulled from $GOCACHE_API_TOKEN
}

resource "gocache_domain" "example" {
  domain = "example.com.br" 
  waf_status = true
  cdn_mode = "ns"
}


resource "gocache_record" "www" {
  domain = gocache_domain.example.domain
  name = "www"
  type = "A"
  proxied = true
  ttl = 60
  content = "192.168.200.100"
}

resource "gocache_record" "teste" {
  domain = gocache_domain.example.domain
  name = "@"
  type = "A"
  proxied = true
  ttl = 60
  content = "192.168.200.101"
}

resource "gocache_record" "bolo" {
  domain = gocache_domain.example.domain
  name = "test"
  type = "CNAME"
  proxied = false
  ttl = 60
  content = "www.example.com.br"
}

output "example_output"{
  value = gocache_domain.example 
}
