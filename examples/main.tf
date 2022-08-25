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


resource "gocache_smart_rules_settings" "mypolicies" {
  domain = gocache_domain.example.domain

  smart_rule {
    match {
      http_version = [ "HTTP/1.0", "HTTP/1.1" ]
    }

    action {
      set_uri = "/invalid"
    }
    metadata {
      status = true
    } 
  }

  smart_rule {
    match {
      request_uri = "/api/*"
    }

    action {
      cache_mode = "full"
      set_request_headers = {"from-api":"1"}
    }

    metadata {
      status = true
    } 
  }

}


resource "gocache_smart_rules_rewrite" "rewrite" {
  domain = gocache_domain.example.domain

  smart_rule {
    match {
      request_uri = "/test2"
    }

    action {
      redirect_type = "301"
      redirect_to = "/index.html"
    }

    metadata {
      status = true
    } 
  }
}

resource "gocache_smart_rules_ratelimit" "limiting" {
  domain = gocache_domain.example.domain

  smart_rule {
    match {
      request_uri = "/api/*"
    }

    action {
      rate_limit_action = "block"
      rate_limit_block_ttl = 86400
      rate_limit_amount = 100
      rate_limit_period = 86400
    }

    metadata {
      status = true
      name = "blocking api"
    } 
  }
}
