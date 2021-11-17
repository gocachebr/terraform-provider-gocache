terraform {
  required_providers {
    gocache = {
      source = "gocache.com.br/gocachebr/gocache"
    }
  }
}

provider "gocache" {
  api_token = "5e28d7603418957c"
}

resource "gocache_domain" "example" {
  domain = "fsantossite.xyz" 
}

resource "gocache_smart_rules_settings" "mypolicies" {
  domain = "fsantossite.xyz"

  smart_rule {
    match {
      header="aa"
      cookie="bb"
      http_version = [ "HTTP/1.0", "HTTP/1.1" ]
    }

    action {
      # set_response_header = "hey:bro"
      set_uri = "/ab"
    }

    metadata {
      status = true
    } 
  }

  smart_rule {
    match {
      request_uri = "/*"
    }

    action {
      cache_mode = "full"
      set_response_headers = {"header":"value","header2":"val"}
      set_request_headers = {"header":"value","ooo":"noo"}
    }

    metadata {
      status = true
    } 
  }

}

resource "gocache_smart_rules_rewrite" "rewrite" {
  domain = "fsantossite.xyz"

  smart_rule {
    match {
      request_uri = "/test2"
    }

    action {
      redirect_type = "301"
      redirect_to = "/yo/ye"
    }

    metadata {
      status = true
    } 
  }

}

resource "gocache_smart_rules_firewall" "firewall" {
  domain = "fsantossite.xyz"

  smart_rule {
    match {
      request_uri = "/test4"
    }

    action {
      firewall_action = "block"
    }

    metadata {
      status = true
    } 
  }

}


# resource "gocache_record" "www" {
#   domain = gocache_domain.example.domain
#   name = "www"
#   type = "A"
#   proxied = true
#   ttl = 60
#   content = "192.168.200.100"
# }

# resource "gocache_record" "teste" {
#   domain = gocache_domain.example.domain
#   name = "@"
#   type = "A"
#   proxied = true
#   ttl = 60
#   content = "192.168.200.101"
# }

# resource "gocache_record" "bolo" {
#   domain = gocache_domain.example.domain
#   name = "test"
#   type = "CNAME"
#   proxied = false
#   ttl = 60
#   content = "www.example.com.br"
# }

# resource "gocache_ssl_certificate" "blab" {
#   domain = gocache_domain.example.domain
#   certificate = file("cert.cert")
#   private_key = file("key.key")
#   enabled = false
# }
# data "gocache_ip_ranges" "gocache" {}

# output "example_output"{
#   value = data.gocache_ip_ranges.gocache
# }


#output "example_output"{
#  value = gocache_ssl_certificate.blab
#}
