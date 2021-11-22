# gocache_domain_dnssec

This resource enables DNSSEC for a domain using GoCache name servers, and exports the attributes needed to configure DNSSEC correctly on its respective registrar.

## Example Usage

Here is an example of a custom certificate configuration

```hcl
resource "gocache_domain" "example" {
  domain = "example.com"
  dns_mode = "ns"
}

resource "gocache_domain_dnssec" "example" {
  domain = "example.com"
}
```

## Argument Reference

This resource supports these arguments:

* `domain` - (Required) The domain to enable DNSSEC.

## Attribute Reference

The following computed attributes are exported:

* `digest` - Digest used for DNSSEC.
* `digest_type` - Digest type id used for DNSSEC.
* `digest_type_name` - Hash algorithm name represented by `digest_type`.
* `algorithm` - DNSSEC algorithm used. 
* `public_key` - Public Key used for DNSSEC.
* `key_tag` - Key tag used for DNSSEC.