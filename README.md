# GoCache Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

## Building The Provider 

Clone repository to: `$GOPATH/src/github.com/gocache/terraform-provider-gocache`

Once the local repository is present, change into that directory and run `make
build-dev`. This will create a new binary in the same directory which will
be loaded for your Terraform operations.

### Using the customer provider locally

Create the following file in your home directory. Note: This file can live
anywhere and isn't restricted to your home directory if you would like it
elsewhere.

```
provider_installation {
  dev_overrides {
    "gocache/gocache" = "<GOPATH>/src/github.com/gocache/terraform-provider-gocache"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

You will need to replace `<GOPATH>` with the **full path** to your GOPATH where
the repository lives, no `~` shorthand.

Once this is in place, you can run your Terraform operations prefixed with
`TF_CLI_CONFIG_FILE=/path/to/the/config/file` and it will load in your custom
overrides. Full details can be found on the [Terraform CLI guide](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine (version 1.15+ is *required*). You'll also need to
correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well
as adding `$GOPATH/bin` to your `$PATH`.

See above for which option suits your workflow for building the provider.

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Managing dependencies

Terraform providers use [Go modules](https://github.com/golang/go/wiki/Modules) to manage the
dependencies. To add or update a dependency, you would run the
following (`v1.2.3` of `foo` is a new package we want to add):

```
$ go get foo@v1.2.3
$ go mod tidy
```

Stepping through the above commands:

- `go get foo@v1.2.3` fetches version `v1.2.3` from the source (if
    needed) and adds it to the `go.mod` file for use.
- `go mod tidy` cleans up any dangling dependencies or references that
  aren't defined in your module file.

(The example above will also work if you'd like to upgrade to `v1.2.3`)

If you wish to remove a dependency, you can remove the reference from
`go.mod` and use the same commands above but omit the initial `go get`.
