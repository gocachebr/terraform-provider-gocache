# gocache_smart_rules_ratelimit

Provides a list of type `ratelimit` Smart Rules for a domain, which effectively are just rate limit rules. Smart Rules are the GoCache platform for transforming the request behaviour under selected conditions and the type `ratelimit` are like the other smart rules with the very same matching criteria just changing the actions to define rate limiting actions. See more in [Rate Limit documentation](https://docs.gocache.com.br/ratelimit/).

## Example Usage

Basic usage:

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
}

resource "gocache_smart_rules_ratelimit" "myrates" {

  domain = gocache_domain.example.domain
  # First rule to block abuse on admin pages
  smart_rule {
    match {
      request_uri = "/admin/*"
    }

    action {
      rate_limit_action = "block"
      rate_limit_block_ttl = 86400
      rate_limit_amount = 100
      rate_limit_period = 86400
    }

    metadata {
      status = true
      name = "blocking admin abuse"
    } 
  }
  # Second rule to avoid api abuse
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
      name = "blocking api abuse"
    } 
  }
}
```

## Argument Reference

The `gocache_smart_rules_ratelimit` argument layout is a complex structure composed of several sub-resources - these resources are laid out below.

### Top-Level Arguments

* `domain` - (Required) The domain where you will apply the policies.
* `smart_rule` - (Required) Each `smart_rule` block sets a rule. More than one block is allowed. The list has no particular order in contrast to other smartrules, each of those rules defined in the block are evaluated.

### Smart Rule Level Arguments

* `match` - (Required) Request criteria needed to trigger the rule action. At least one option is needed. See the options in [match level arguments](match-level-arguments).
* `action` - (Required) Actions applied to a request when the criteria is matched. At least one option is needed. See the options in [action level arguments](action-level-arguments).
* `metadata` - (Required) Rule metadata. See the options in [metadata level arguments](metadata-level-arguments).

### Match Level Arguments

Some fields can use regex patterns. To know more about regex, see the documentation in [Smart Rules Regexp](https://docs.gocache.com.br/smart_rules/smart_rules-regexp/).

* `request_method` - (Optional) List of HTTP request methods matched by the rule. Possible values `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`.
* `scheme` - (Optional) List of HTTP request scheme matched by the rule. Possible values `http`, `https` or `http*`
* `hostname` - (Optional) HTTP request hostname matched by the rule. For all hostnames, set `*`.
* `request_uri` - (Optional) HTTP request URI pattern matched by the rule. You can use regex in this field.
* `http_user_agent` - (Optional) Request user agent matched by the rule. You can use regex in this field.
* `cookie` - (Optional) List of request cookie names matched by the rule. This field supports wildcard.
* `cookie_content` - (Optional) Cookie values for a cookie matched by the rule. You need to declare cookie and its value separated by `=`, like `wp_test=yes`. You can use regex in the value.
* `http_referer` - (Optional) HTTP request referer matched by the rule.You can use regex in this field.
* `remote_address` - (Optional) Request client IP ou IP ranges matched by the rule. Only `/32`, `/24` and `/16` ranges are allowed.
* `http_version` - (Optional) List of request HTTP version matched by the rule. Possible values are `HTTP/1.0`,`HTTP/1.1`,`HTTP/2.0`.
* `header` - (Optional) HTTP headers matched by the rule. Header name and value are separated by `:`. You can use regex in the value.
* `origin_country` - (Optional) List of origin countries matched by the rule in the ISO format, like `BR`.
* `origin_continent` - (Optional) List of origin continents matched by the rule in the ISO format. Possible values `SA`, `NA`, `OC`, `EU`, `AF`, `AS`.
* `device_type` - (Optional) list of device types based on user agent matched by the rule. Possible values `mobile`,`desktop`,`bot`,`na`
* `bots` - (Optional) If the bot is legitimate. Possible values `known`, `others`

### Action Level Arguments

* `rate_limit_action` - (Required) When `block`, the matched request is denied with status `429`, when `challenge`, a captcha challenge is returned and when `simulate`, the request is allowed, even when WAF blocks it. Possible values: `block`, `challenge`, `simulate`
* `rate_limit_amount` - (Required) Amount of requests from a single ip needed for the action to be taken.
* `rate_limit_period` - (Required) Defines the minimum time window (seconds) to the requests reach `rate_limit_amount` for the action can be taken. Possible values `86400`, `43200`, `14400`, `3600`, `1800`, `600`, `300`, `120`, `60`, `30`, `10`
* `rate_limit_block_ttl` - (Required) Defines the duration of the blockage (in seconds) after the action has been taken. Possible values: `86400`, `43200`, `14400`, `3600`, `1800`, `600`, `300`, `120`, `60`

### Metadata Level Arguments

* `status` - (Required) If `true`, the rule is activated.
* `name` - (Required) Rule name, can be any text like: `my custom rule to block bots`

## Import

Domain resource can be imported using the domain address.

```
$ terraform import gocache_smart_rules_ratelimit.myrates example.com.br
```