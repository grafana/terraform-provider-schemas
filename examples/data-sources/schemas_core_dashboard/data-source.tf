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
  field_config = {
    defaults = {

    }
    overrides = []
  }
  options = {

  }
  panel_options = {
    content = "# hello world"

  }
  transformations = [

  ]
  code_options = {

  }
  grid_pos = {
    h = 3
    w = 24
    x = 0
    y = 0
  }

  title = "test title"

  type      = "text"
  text_mode = "markdown"
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
