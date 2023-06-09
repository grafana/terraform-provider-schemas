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
  config_json = data.schemas_core_dashboard.prometheus_dashboard.rendered_json
}

resource "grafana_data_source" "prometheus" {
  type               = "prometheus"
  name               = "Prometheus"
  url                = "http://localhost:9090"
  basic_auth_enabled = false
}

data "schemas_core_dashboard" "prometheus_dashboard" {
  title = "Prometheus Dashboard"

  time = {
    from = "now-1h"
  }

  panels = [
    data.schemas_panel_stat.disk.rendered_json,
    data.schemas_panel_time_series.cpu_usage.rendered_json,
    data.schemas_panel_time_series.disk_usage.rendered_json,
  ]
}

data "schemas_panel_stat" "disk" {
  title = "Disk"

  datasource = {
    uid  = resource.grafana_data_source.prometheus.uid
    type = resource.grafana_data_source.prometheus.type
  }

  grid_pos = {
    h = 4
    w = 6
    x = 8
    y = 0
  }

  field_config = {
    defaults = {
      unit = "percentunit"
    }
  }

  targets = [
    data.schemas_query_prometheus.disk_target.rendered_json
  ]
}

data "schemas_query_prometheus" "cores_target" {
  expr   = "count(count by(cpu)(node_cpu_seconds_total))"
  format = "time_series"
  ref_id = "A"
}

data "schemas_panel_time_series" "cpu_usage" {
  title = "CPU usage"

  datasource = {
    uid  = resource.grafana_data_source.prometheus.uid
    type = resource.grafana_data_source.prometheus.type
  }

  grid_pos = {
    h = 8
    w = 12
    x = 0
    y = 4
  }

  options = {
    legend = {
      calcs = [
        "mean",
        "lastNotNull",
        "max",
        "min"
      ]
      display_mode = "table"
      placement    = "bottom"
      show_legend  = true
    }
  }

  targets = [
    data.schemas_query_prometheus.cpu_target.rendered_json
  ]
}

data "schemas_query_prometheus" "cpu_target" {
  expr   = "sum by(mode) (irate(node_cpu_seconds_total{mode!=\"idle\"}[5m]))"
  format = "time_series"
  ref_id = "A"
}

data "schemas_panel_time_series" "disk_usage" {
  title = "Disk I/O utilization"

  datasource = {
    uid  = resource.grafana_data_source.prometheus.uid
    type = resource.grafana_data_source.prometheus.type
  }

  grid_pos = {
    h = 8
    w = 12
    x = 12
    y = 4
  }

  options = {
    legend = {
      calcs = [
        "mean",
        "lastNotNull",
        "max",
        "min"
      ]
      display_mode = "table"
      placement    = "bottom"
      show_legend  = true
    }
  }

  field_config = {
    defaults = {
      unit = "percentunit"
      min  = 0
      max  = 1
    }
  }

  targets = [
    data.schemas_query_prometheus.disk_target.rendered_json
  ]
}

data "schemas_query_prometheus" "disk_target" {
  expr   = "sum(irate(node_disk_io_time_seconds_total[5m]))"
  format = "time_series"
  ref_id = "A"
}
