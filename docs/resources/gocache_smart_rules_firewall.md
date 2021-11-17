# gocache_smart_rules_firewall

Provides a list of type `firewall` Smart Rules for a domain. Smart Rules are the GoCache platform for transforming the request behavior under selected conditions. Smart Rules of type `firewall` are created to apply firewall actions on the matched conditions. Only one resource per domain can be used. See more in https://docs.gocache.com.br/smart_rules/.

## Example Usage

Basic usage:

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
}

resource "gocache_smart_rules_firewall" "mypolicies" {
  domain = "example.com"

  # Smart rule with precedence 0
  smart_rule {
    match {
      request_uri = "/admin/*"
      origin_country = ["BR"]
    }
    action {
      firewall_action = "accept"
    }
    metadata {
      notes = "This rule acts along with the rule bellow, allowing only requests for /admin/* originated from Brasil"
    }
  }

  # Smart rule with precedence 1
  smart_rule {
    match {
      request_uri = "/admin/*"
    }
    action {
      firewall_action = "block"
    }
    metadata {
      notes = "This rule acts along with the rule above, blocking only all requests for /admin/* that is not originated from Brasil"
    }
  }
}
```

## Argument Reference

The `gocache_smart_rules_firewall` argument layout is a complex structure composed of several sub-resources - these resources are laid out below.

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
* `scheme` - (Optional) List of HTTP request scheme matched by the rule. Possible values `http`, `https` or `http*`
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

* `firewall_action` - (Required) When `block`, the matched request are denied with status `403`, when `challenge`, a captcha challenge is returned and when `accept`, the request is allowed, even when WAF blocks it.

### Metadata Level Arguments

* `status` - (Optional) If `true`, the rule is activated.
* `notes` - (Optional) Smart Rule description

## Import

Domain resource can be imported using the domain address.

```
$ terraform import gocache_smart_rules_firewall.mypolicies example.com.br
```