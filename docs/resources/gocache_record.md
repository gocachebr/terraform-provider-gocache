# gocache_record

Use the `gocache_record` resource to configure a DNS record for your domain, if the domain is set on DNS mode, or to configure a CNAME entry point for your record, if the domain is set on CNAME mode. See the `gocache_domain` resource documentation to understand about the `dns_mode` argument. The `gocache_domain` resource needs to be created before `gocache_record`.

## Example Usage

Here is an example of a configuration for a domain in `ns` mode. All DNS records needs to be declared.

```hcl
resource "gocache_domain" "example" {
  domain = "example.com.br"
  dns_mode = "ns"
}

resource "gocache_record" "root" {
  domain = gocache_domain.example.domain
  name = "@"
  type = "A"
  content = "1.2.3.4"
  ttl = 3600
  proxied = true
}

resource "gocache_record" "www" {
  domain = gocache_domain.example.domain
  name = "www"
  type = "CNAME"
  content = "example.com.br"
  ttl = 3600
  proxied = true
}

resource "gocache_record" "mail" {
  domain = gocache_domain.example.domain
  name = "mail"
  type = "MX"
  content = "mail-server.com"
  ttl = 3600
  priority = 10
}

resource "gocache_record" "sip" {
  domain = gocache_domain.example.domain
  name = "_sip._tcp"
  type = "SRV"
  content = "sipserver.example.com"
  ttl = 3600
  priority = 10
  weight = 5
  port = 5060
}
```

Here is an example of a configuration for a domain in `cname` mode. Only records representing subdomains that are going to be proxied by GoCache needs to be declared.

```hcl
resource "gocache_domain" "example" {
  domain = "example.com.br"
  dns_mode = "cname"
}

resource "gocache_record" "www" {
  domain = gocache_domain.example.domain
  name = "www"
  content = "example-backend.com.br"
}

resource "gocache_record" "api" {
  domain = gocache_domain.example.domain
  name = "api"
  content = "example-api-backend.com.br"
}
```

## Argument Reference

This resource supports these arguments for domains on both `ns` and `cname` modes:

* `domain` - (Required) The domain to add the record to.
* `name` - (Required) The subdomain that the record will set. To declare an `A` record for the root domain on `ns` mode, set `@`. 
* `content` - (Required) The value the record returns. When `dns_mode` is `ns` and record `type` is not `TXT` or `dns_mode` is `cname`, it indicates a server address where the service of that subdomain is hosted. `A` type records only accepts IPs while `CNAME` and `SRV` type records only accepts hostnames. `TXT` type records is used to returns generic text.

## Specific arguments for `dns` mode

This section lists specific arguments when `dns_mode` is set to `ns`

* `type` - (Required) The DNS record type. Possible values are `A`, `CNAME`, `NS`, `TXT`, `MX`, `SRV`.
* `ttl` - (Required) The DNS record type. Possible values are `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `43200`, `86400`.

### Arguments required by `A` and `CNAME` records

The following argument is accepted and required only by `A` and `CNAME` records

* `proxied` - (Required) When set to `true`, the subdomain traffic is proxied through GoCache, enabling the most of the services. When set to `false`, the subdomain traffic is not delivered by GoCache.

### Optional arguments accepted by `MX` and `SRV` records

* `priority` - (Optional) When two or more records are declared for the same subdomain, the records with the lowest value is prioritized

### Arguments required by `SRV` records

* `port` - (Required) The port number the service configured by this record listen
* `weight` - (Required) When two or more records are declared for the same subdomain, with the same `priority` argument value or without it, the traffic is distributed proportionaly based on the value of this field

## Attribute Reference

The following computed attributes are exported:

* `record_id` - The record ID in GoCache

## Import

Record resource can be imported using a composite ID formed of domain and its record ID, as returned from [API](https://docs.gocache.com.br/api/#api-Websites_e_DNS-GetDNS)

```
$ terraform import gocache_record.www example.com.br/99990
```