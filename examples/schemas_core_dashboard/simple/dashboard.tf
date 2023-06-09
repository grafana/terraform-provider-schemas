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
  config_json = data.schemas_core_dashboard.dashboard.rendered_json
}

data "schemas_core_dashboard" "dashboard" {
  title         = "Faro dashboard"
  uid           = "faro-terraform-demo"
  description   = "Dashboard for Faro"
  graph_tooltip = 1

  panels = [
    data.schemas_panel_time_series.requests.rendered_json,
  ]
}

data "schemas_panel_time_series" "requests" {
  title = "Requests / sec"

  datasource = {
    uid  = "ops-cortex"
    type = "prometheus"
  }

  targets = [
    data.schemas_query_prometheus.requests_target.rendered_json
  ]

  field_config = {
    defaults = {
      unit = "reqps"
    }
  }

  grid_pos = {
    h = 8
    w = 24
  }
}

data "schemas_query_prometheus" "requests_target" {
  expr   = "sum by (status_code) (rate(request_duration_seconds_count{job=~\".*/faro-api\"}[$__rate_interval]))"
  format = "time_series"
  ref_id = "A"
}
