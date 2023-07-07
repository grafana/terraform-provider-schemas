# Grafana Schemas Terraform Provider (experimental)

This is a generated provider to manage Grafana dashboards. 

The code in this repository should be considered experimental. It comes with no support, but we are keen to receive feedback on the product and suggestions on how to improve it, though we cannot commit to resolution of any particular issue. No SLAs are available. It is not meant to be used in production environments, and the risks are unknown/high. 

Our goal is for these generated data sources to become a part of our official Grafana provider once this project becomes more mature. At this time, we are not planning to create a migration path for data sources created with the `schemas` provider when they are merged into the `grafana` provider. Also, this work is subject to signficant changes as we iterate towards that level of quality.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0

## Maturity

Grafana Labs defines experimental features as follows:

> Projects and features in the Experimental stage are supported only by the Engineering
teams; on-call support is not available. Documentation is either limited or not provided
outside of code comments. No SLA is provided.
>
> Experimental projects or features are primarily intended for open source engineers who
want to participate in ensuring systems stability, and to gain consensus and approval
for open source governance projects.
>
> Projects and features in the Experimental phase are not meant to be used in production
environments, and the risks are unknown/high.

## Usage

This provider should be used with the [Grafana Terraform Provider](https://registry.terraform.io/providers/grafana/grafana/latest). It is generated from the [CUE schemas](https://github.com/grafana/grafana/blob/main/kinds/dashboard/dashboard_kind.cue) defined in the Grafana repository to ensure it stays up-to-date.

Configure the [Grafana Terraform Provider](https://registry.terraform.io/providers/grafana/grafana/latest) to access your Grafana instance (see [documentation](https://grafana.com/docs/grafana-cloud/infrastructure-as-code/terraform/)). Then use this Grafana Schemas Terraform provider to create your dashboards.

```
terraform {
  required_providers {
    schemas = {
      source  = "grafana/schemas"
      version = "0.1.0"
    }
    grafana = {
      source  = "grafana/grafana"
      version = "1.36.1"
    }
  }
}

provider "grafana" {
  url  = "http://localhost:3000"
  auth = "admin:admin"
}

data "schemas_core_dashboard" "my_dashboard" {
  title = "My Dashboard"

  time = {
    from = "now-1h"
  }

  panels = [
    data.schemas_panel_stat.my_panel.rendered_json,
  ]
}

data "schemas_panel_stat" "my_panel" {
  title = "My Panel"

  grid_pos = {
    h = 4
    w = 6
    x = 0
    y = 0
  }
}

resource "grafana_dashboard" "example" {
  config_json = data.schemas_core_dashboard.my_dashboard.rendered_json
}
```

More examples are available in the ["examples"](https://github.com/grafana/terraform-provider-schemas/tree/main/examples/data-sources/schemas_core_dashboard) folder.

## Development

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

### Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

If you want to re-generate the provider, use the `make generate` command.

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

### Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
