# 0.4.0 (Unreleased)

# 0.3.0 (November 22th, 2021)

FEATURES:

* **New Resource:** `gocache_domain_dnssec`
* **New Resource:** `gocache_smart_rules_settings` ([#1](https://github.com/gocachebr/terraform-provider-gocache/issues/1))
* **New Resource:** `gocache_smart_rules_rewrite` ([#1](https://github.com/gocachebr/terraform-provider-gocache/issues/1))
* **New Resource:** `gocache_smart_rules_firewall` ([#1](https://github.com/gocachebr/terraform-provider-gocache/issues/1))

BUG FIXES:

* resource/gocache_domain: Fixed bad behaviour when adding a domain that already exists

# 0.2.2 (September 9th, 2021)

IMPROVEMENTS:

* `resource/gocache_record`: Improvements on scalability for domains with many records

# 0.2.1 (September 2nd, 2021)

BUG FIXES:

* resource/gocache_ssl_certificate: Fixed missing documentation
* data-source/gocache_ip_ranges: Fixed missing documentation

# 0.2.0 (September 2nd, 2021)

FEATURES:

* **New Resource:** `gocache_ssl_certificate`
* **New Data Source:** `gocache_ip_ranges`

## 0.1.2 (August 31th, 2021)

BUG FIXES:

* `resource/gocache_domain`: Fixed inexpected behavior

## 0.1.1 (August 30th, 2021)

BUG FIXES:

* `provider/gocache`: Fixed crash in authentication

## 0.1.0 (August 27th, 2021)

FEATURES:

* **New Resource:** `gocache_domain`
* **New Resource:** `gocache_record`