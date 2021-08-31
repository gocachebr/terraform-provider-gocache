# GoCache Provider

The GoCache provider is used to interact with the Content Delivery Network (CDN), Web Application Firewall and other network and security services provided by GoCache.

In order to use this Provider, you must have an active account with GoCache. If you don't have an account you can signup on the service with a 7 days trial period on https://painel.gocache.com.br/trial and you can see the pricing information on the page https://www.gocache.com.br/planos/.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the GoCache Provider
provider "gocache" {
  api_token = "yourtoken"
}

# Create a Domain
resource "gocache_domain" "mydomain" {
  # ...
}

# Create a subdomain record
resource "gocache_record" "mysubdomain" {
  # ...
}

# Create cache, security and routing policies
resource "gocache_smart_rules" "mypolicies" {
  # ...
}
```

## Authentication

The methods for providing credentials for the GoCache provider is based on an API token. Two methods are supported:

- Static API token
- Environment variables


### Static API Key

Static credentials can be provided by adding a `api_token` in-line in the GoCache provider block:

Usage:

```hcl
provider "gocache" {
  api_token = "test"
}

resource "gocache_domain" "mydomain" {
  # ...
}
```

See how to get your api token in order to use the Provider on https://docs.gocache.com.br/painel/minha_conta/minha_conta-conta/.


### Environment variables

You can provide your API token via `GOCACHE_API_TOKEN` environment variable, representing your GoCache API token.

```hcl
provider "gocache" {
  # token pulled from $GOCACHE_API_TOKEN
}

resource "gocache_domain" "mydomain" {
  # ...
}
```

Usage:

```
$ export GOCACHE_API_TOKEN="yourtoken"
$ terraform plan
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `api_token` - (Optional) This is the API token. It must be provided, but it can also be sourced from the `GOCACHE_API_TOKEN` environment variable