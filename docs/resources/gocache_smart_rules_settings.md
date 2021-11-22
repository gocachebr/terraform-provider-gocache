# gocache_smart_rules_settings

Provides a list of type `settings` Smart Rules for a domain. Smart Rules are the GoCache platform for transforming the request behavior under selected conditions. Smart Rules of type `settings` transform the behaviour in terms of cache, routing and status and configuration of other GoCache products. Only one resource per domain can be used. See more in https://docs.gocache.com.br/smart_rules/.

## Example Usage

Basic usage:

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
}

resource "gocache_smart_rules_settings" "mypolicies" {
  domain = "example.com"

  # Smart rule with precedence 0
  smart_rule {
    match {
      hostname = "assets.gocache.com.br"
    }
    action {
      backend = "mybucket.s3.amazon.com"
      cache_mode = "full"
    }
    metadata {
      notes = "Smart rule to cache and route static content to bucket"
    }
  }

  # Smart rule with precedence 1
  smart_rule {
    match {
      hostname = "api.gocache.com.br"
      request_uri = "/v2/*"
    }
    action {
      backend = "mynewapibackend.com"
      set_uri = "/$1"
    }
    metadata {
      notes = "Smart rule to route API v2 to new backend"
    }
  }
}
```

## Argument Reference

The `gocache_smart_rules_settings` argument layout is a complex structure composed of several sub-resources - these resources are laid out below.

### Top-Level Arguments

* `domain` - (Required) The domain where you will apply the policies.
* `smart_rule` - (Required) Each `smart_rule` block sets a rule. More than one block is allowed. List from top to bottom in order of precedence. The topmost block will have precedence 0. See below the structure of each block.

### Smart Rule Level Arguments

* `match` - (Required) Request criteria needed to trigger the rule action. At least one option is needed. See the options in [match level arguments](match-level-arguments).
* `action` - (Required) Actions applied to a request when the criteria is matched. At least one option is needed. See the options in [action level arguments](action-level-arguments).
* `metadata` - (Optional) Rule metadata. See the options in [metadata level arguments](metadata-level-arguments).

### Match Level Arguments

Some fields can use regex patterns. To know more about regex, see the documentation in https://docs.gocache.com.br/smart_rules/smart_rules-regexp/.

* `request_method` - (Optional) List of HTTP request methods matched by the rule. Possible values `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`.
* `scheme` - (Optional) List of HTTP request scheme matched by the rule. Possible values `http`, `https` or `http*`.
* `hostname` - (Optional) HTTP request hostname matched by the rule. For all hostnames, set `*`.
* `request_uri` - (Optional) HTTP request URI pattern matched by the rule. You can use regex in this field.
* `http_user_agent` - (Optional) Request user agent matched by the rule. You can use regex in this field.
* `cookie` - (Optional) List of cequest cookie names matched by the rule. This field supports wildcard.
* `cookie_content` - (Optional) Cookie values for a cookie matched by the rule. You need to declare cookie and its value separeted by `=`, like `wp_test=yes`. You can use regex in the value.
* `http_referer` - (Optional) HTTP request referer matched by the rule.You can use regex in this field.
* `remote_address` - (Optional) Request client IP ou IP ranges matched by the rule. Only `/32`, `/24` and `/16` ranges are allowed.
* `http_version` - (Optional) List of request HTTP version matched by the rule. Possible values are `HTTP/1.0`,`HTTP/1.1`,`HTTP/2.0`.
* `header` - (Optional) HTTP headers matched by the rule. Header name and value are separated by `:`. You can use regex in the value.
* `origin_country` - (Optional) List of origin countries matched by the rule in the ISO format, like `BR`.
* `origin_continent` - (Optional) List of origin continents matched by the rule in the ISO format. Possible values `SA`, `NA`, `OC`, `EU`, `AF`, `AS`.
* `device_type` - (Optional) list of device types based on user agent matched by the rule. Possible values `mobile`,`desktop`,`bot`,`na`
* `bots` - (Optional) If the bot is legitimate. Possible values `known`, `others`

### Action Level Arguments

* `set_host` - (Optional) Overrides the header `Host` when requesting the origin.
* `backend` - (Optional) Sets a different backend as origin for the request. It can be an IP address or a hostname.
* `set_uri` - (Optional) Overrides the URI when requesting the origin. If there is one or more wildcard in the `request_uri` match criteria, you can use variables like `$1` to pass the respective string sequence on the backend request.
* `custom_cache_key` - (Optional) Customizes cache key. You can use variables like `$cookie_NAME` to group cached requests by specific cookie values, `$http_NAME`, to group by specific header values, `$1` to group by URI patterns respective to `request_uri` wildcards, `$geoip2_data_country_code` to group by request's country and `$geoip2_data_continent_code` to group by request's continent.
* `expires_ttl` - (Optional) Overrides the `expires` header that will be sent with assets requests. Possible values are `-1`, `10`, `30`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `43200`, `86400`, `172800`, `345600`, `604800`, `1296000`, `2592000`, `15552000`, `31536000`.
* `caching_behavior` - (Optional) Overrides the caching behaviour when there are different querystrings on a matched URL. Possible values are `default` and `ignore_query_string`.
* `cache_mode` - (Optional) When `off` no cache are going to be applied, when `default`, only static assets will be cached, when `full`, all requests will be cached.
* `cache_ttl` - (Optional) Overrides the time after the CDN cache will be expired. Possible values are `300`, `600`, `900`, `1800`, `3600`, `7200`, `14400`, `86400`, `172800`, `604800`. Default `86400`..
* `cors` - (Optional) Defines Cross-Origin headers for the matched requests, based on the address informed.
* `ssl_mode` - (Optional) The `full` value indicates that when a client sends an https request, GoCache will comunite with origin server on port 443 and the `partial` value indicates all requests sent to origin server by GoCache will be on port 80..
* `hide_header` - (Optional) Omits a response header informed.
* `encoding_header` - (Optional) Defines the Accept-Encoding header.
* `set_request_headers` - (Optional) Map object of request headers to be added on the upstream request, where header names are the keys. You can use variables like `$cookie_NAME` to group cached requests by specific cookie values, `$http_NAME`, to group by specific header values, `$geoip2_data_country_code` to group by request's country and `$geoip2_data_continent_code` to group by request's continent.
* `set_response_headers` - (Optional) Map object of response headers to be added on the response, where header names are the keys.
* `signed_url_key` - (Optional) Defines a signed URL key to request S3 protected assets.
* `signed_url_type` - (Optional) Defines a signed URL type to request S3 protected assets. Possible values `s3qs`, `off`
* `cache_301` - (Optional) Defines if responses with status 301 will be cache.
* `cache_302` - (Optional) Defines if responses with status 302 will be cache..
* `cache_404` - (Optional) Defines if responses with status 404 will be cache..
* `gzip_status` - (Optional) Activates the GZIP compression on assets server by the CDN.
* `ignore_cache_control` - (Optional) Ignores the `cache-control` header sent by the origin server.
* `ignore_expires` - (Optional) Ignores the `expires` header sent by the origin server.
* `ignore_vary` - (Optional) Ignores the `vary` header sent by the origin server.
* `waf_status` - (Optional) Enables the WAF when status in `true`.
* `waf_mode` - (Optional) Defines the default action on a request when it is detected by WAF. When `simulate`, the request is allowed but logged in security events. When `challenge`, a captcha challenge will be sent to the client. Finally, when 'block', the request is blocked with status 403.
* `waf_level` - (Optional) Defines the sensitivity level of WAF. Possible values are `low`, `medium`, `high`.
* `ratelimit_status` - (Optional) Ignores the request in rate limiting.
* `image_optimize` - (Optional) Enables the image optimizer. This option needs to be enabled for all image optization options bellow work. See the https://www.gocache.com.br/planos/ page to know the princing of this feature.
* `image_optimize_webp` - (Optional) Enables the image conversion to WEBP for browser that support the format.
* `image_optimize_progressive` - (Optional) Enables the image conversion to progressive JPEG for browsers that support the format. Convertion to WEBP takes precedence.
* `image_optimize_metadata` - (Optional) Removes unnecessary metadata from the images.
* `image_optimize_level` - (Optional) Sets the compression level on the image optimization proccess. Choose `0` to maintain the image quality, `90` for low compression, `75` for medium and `65` for high.
* `waf_rule_action` - (Optional) Customizes on or more WAF rules behaviour. Possible behaviours are `DISABLE`, `DEFAULT`, `BLOCK`, `CHALLENGE`, `SIMULATE`. Send the behaviour along with all rules you want this behaviour, separated by `:`, like `CHALLENGE:gocache-v1/90015,gocache-v1/41*`. To send more behaviours, separate by `;`, like `CHALLENGE:gocache-v1/90015;SIMULATE:gocache-v1/41*;BLOCK:gocache-v1/10*;DISABLE:gocache-v1/*`.

### Metadata Level Arguments

* `status` - (Optional) If `true`, the rule is activated.
* `notes` - (Optional) Smart Rule description

## Import

Domain resource can be imported using the domain address.

```
$ terraform import gocache_smart_rules_settings.mypolicies example.com.br
```