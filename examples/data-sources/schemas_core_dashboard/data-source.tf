terraform {
  required_providers {
    schemas = {
      source  = "terraform.local/grafana/schemas"
      version = "1.0.0"
    }
    grafana = {
      source  = "grafana/grafana"
      version = "1.36.1"
    }
  }
}

data "schemas_panel_text" "test" {
  transformations = [
  ]
  options = {
    content = "# test content"
  }
  grid_pos = {
    h = 10
    w = 24
  }


  title = "test title"
}

data "schemas_core_dashboard" "example" {
  uid    = "test"
  title  = "test"
  panels = [data.schemas_panel_text.test.to_json]

}

resource "local_file" "foo" {
  content  = data.schemas_core_dashboard.example.to_json
  filename = "${path.module}/foo.json"
}

provider "grafana" {
  url  = "http://localhost:3000"
  auth = "admin:admin"
}

resource "grafana_dashboard" "example" {
  config_json = data.schemas_core_dashboard.example.to_json
}
