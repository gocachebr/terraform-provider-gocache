# gocache_ip_ranges

Use this data source to get the IP ranges of GoCache edge nodes. You can use the IPs to add in a firewall whitelist on your origin infrastructure for example. IPV6 support are going to be released soon.

## Example Usage

Here is an example of a custom certificate configuration

```hcl
data "gocache_ip_ranges" "gocache" {}

resource "aws_security_group" "from_gocache" {
  name = "from_gocache"

  ingress {
    from_port         = "443"
    to_port           = "443"
    protocol          = "tcp"
    cidr_blocks       = data.gocache_ip_ranges.gocache.ipv4_cidr_blocks
  }
}
```

## Attributes Reference

* `cidr_blocks` - The lexically ordered list of all CIDR blocks.
* `ipv4_cidr_blocks` - The lexically ordered list of only the IPv4 CIDR blocks.