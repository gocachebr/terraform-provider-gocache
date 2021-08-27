# gocache_domain

Provides a domain in your account in GoCache, with all the domain related configurations. In order to effectively start to use GoCache, you need to declare at least one `gocache_record` block to configure an entry point (subdomain) and its backend address, and after, make the appropriate DNS changes.

## Example Usage

Basic usage if you are going to use the GoCache name servers on the domain

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
  dns_mode = "ns"
}
```

Basic usage if you are going to keep your name servers

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
  dns_mode = "cname"
}
```

Basic usage if you are going to use the GoCache name servers and activate WAF in the low level and simulate mode

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
  dns_mode = "ns"
  waf_status = true
  waf_level = "low"
  waf_mode = "simulate"
}
```

## Argument Reference

The following arguments are supported in the `gocache_domain` resource:

* `domain` - (Required) The domain you want to add.
* `cache_ttl` - (Optional) The time after the CDN cache will be expired. Possible values are `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `86400`, `172800`, `604800`. Default `86400`.
* `deploy_mode` - (Optional) Deactivates cache for debbugging purpose. Default `false`.
* `smart_cache_status` - (Optional) Activates Smart Cache, a template of cache rules for dinamic cache on popular Content Management Systems. Default `false`.
* `smart_cache_template` - (Optional) Selects the Smart Cache template. Possible values are `custom`, `wordpress`, `magento`, `joomla`. Default `custom`.
* `smart_cache_ttl` - (Optional) The time after the cache from Smart Cache will be expired. Possible values are `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `86400`, `172800`, `604800`. Default `14400`.
* `dns_mode` - (Optional) Choose `ns` to use the GoCache name servers or choose `cname` if you are goint to maintain your current name servers. Default `ns`.
* `gzip_status` - (Optional) Activates the GZIP compression on assets server by the CDN. Default `true`.
* `expires_ttl` - (Optional) The define the `expires` header that will be sent with assets requests to determines the time they will be cached in browsers. Possible values are `-1`, `10`, `30`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `43200`, `86400`, `172800`, `345600`, `604800`, `1296000`, `2592000`, `15552000`, `31536000`. Default `14400`.
* `ignore_vary` - (Optional) Ignores the `vary` header sent by the origin server. Default `true`.
* `ignore_cache_control` - (Optional) Ignores the `cache-control` header sent by the origin server. Default `true`.
* `ignore_expires` - (Optional) Ignores the `expires` header sent by the origin server. Default `true`.
* `ssl_mode` - (Optional) The `full` value indicates that when a client sends an https request, GoCache will comunite with origin server on port 443 and the `partial` value indicates all requests sent to origin server by GoCache will be on port 80. Default `full`.
* `cache_301` - (Optional) Defines if responses with status 301 will be cache. Default `false`.
* `cache_302` - (Optional) Defines if responses with status 302 will be cache. Default `false`.
* `cache_404` - (Optional) Defines if responses with status 404 will be cache. Default `false`.
* `header_device_type` - (Optional) If activated, a header `GoCache-Device-Type` will be sent to origin server containing the device type that made the request. Default `false`.
* `header_geoip_continent` - (Optional) If activated, a header `GoCache-GeoIP-Continent` will be sent to origin server containing the continent where the request was from in ISO-ALPHA-2 notation. Default `false`.
* `header_geoip_country` - (Optional) If activated, a header `GoCache-GeoIP-Country` will be sent to origin server containing the country where the request was from in ISO-ALPHA-2 notation. Default `false`.
* `header_geoip_org` - (Optional) If activated, a header `GoCache-GeoIP-Org` will be sent to origin server containing the Autonomous System Number (ASN) and organization responsible for the IP that made the request. Default `false`.
* `caching_behavior` - (Optional) When `default`, each querystring on an URL will generate another cache, when `ignore` all querystrings in the same URL will be served by the same cache. Default `default`.
* `waf_status` - (Optional) Enables the WAF when status in `true`. Default `false`.
* `waf_level` - (Optional) Defines the sensitivity level of WAF. Possible values are `low`, `medium`, `high`. Default `low`.
* `waf_mode` - (Optional) Defines the default action on a request when it is detected by WAF. When `simulate`, the request is allowed but logged in security events. When `challenge`, a captcha challenge will be sent to the client. Finally, when 'block', the request is blocked with status 403. Default `simulate`.
* `expires_bypass_sec` - (Optional) Sets the expiration time of the cookie sent for clients that was approved on challenge. Possible values are `-1`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `43200`, `86400`, `172800`, `345600`, `604800`, `1296000`, `2592000`, `15552000`, `31536000`. Default `-1`.
* `tls10` - (Optional) Enables the TLS 1.0 protocol. Default `true`.
* `tls11` - (Optional) Enables the TLS 1.1 protocol. Default `true`.
* `image_optimize` - (Optional) Enables the image optimizer. This option needs to be enabled for all image optization options bellow work. See the https://www.gocache.com.br/planos/ page to know the princing of this feature. Default `false`.
* `image_optimize_webp` - (Optional) Enables the image conversion to WEBP for browser that support the format. Default `false`.
* `image_optimize_progressive` - (Optional) Enables the image conversion to progressive JPEG for browsers that support the format. Convertion to WEBP takes precedence. Default `false`.
* `image_optimize_metadata` - (Optional) Removes unnecessary metadata from the images. Default `false`.
* `image_optimize_level` - (Optional) Sets the compression level on the image optimization proccess. Choose `0` to maintain the image quality, `90` for low compression, `75` for medium and `65` for high. Default `false`.
* `rate_limit` - (Optional) Enables the rate limiting. See the https://www.gocache.com.br/planos/ page to know the princing of this feature. Default `false`.
* `rate_limit_ignore_good_bots` - (Optional) Ignores good bots in rate limiting. Default `true`.
* `rate_limit_ignore_static_content` - (Optional) Ignores static content in rate limiting. Default `true`.

## Import

Domain resource can be imported using the domain address.

```
$ terraform import gocache_domain.example example.com.br
```