---
page_title: "Schemas Provider"
description: |-
  Generated provider to manage Grafana dashboards. 
---

# Schemas Provider

This provider provides data sources to manage Grafana [dashboards](https://grafana.com/docs/grafana/latest/dashboards/). You should use it with the [Grafana Terraform Provider](https://registry.terraform.io/providers/grafana/grafana/latest). It is generated from the [CUE schemas](https://github.com/grafana/grafana/blob/main/kinds/dashboard/dashboard_kind.cue) defined in the Grafana repository to ensure it stays up-to-date.

The code in this repository should be considered experimental. It comes with no support, but we are keen to receive feedback on the product and suggestions on how to improve it, though we cannot commit to resolution of any particular issue. No SLAs are available. It is not meant to be used in production environments, and the risks are unknown/high. 

Our goal is for these generated data sources to become a part of our official Grafana provider once this project becomes more mature. At this time, we are not planning to create a migration path for data sources created with the `schemas` provider when they are merged into the `grafana` provider. Also, this work is subject to signficant changes as we iterate towards that level of quality.

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

## Example usage

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
resource "grafana_dashboard" "example" {
  config_json = data.schemas_core_dashboard.my_dashboard.rendered_json
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
```