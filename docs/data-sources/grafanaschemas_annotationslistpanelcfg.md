---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafanaschemas_annotationslistpanelcfg Data Source - terraform-provider-grafana-schemas"
subcategory: ""
description: |-
  TODO description
---

# grafanaschemas_annotationslistpanelcfg (Data Source)

TODO description



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `panel_options` (Attributes) (see [below for nested schema](#nestedatt--panel_options))

### Read-Only

- `to_json` (String) This datasource rendered as JSON

<a id="nestedatt--panel_options"></a>
### Nested Schema for `panel_options`

Required:

- `limit` (Number)
- `navigate_after` (String)
- `navigate_before` (String)
- `navigate_to_panel` (Boolean)
- `only_from_this_dashboard` (Boolean)
- `only_in_time_range` (Boolean)
- `show_tags` (Boolean)
- `show_time` (Boolean)
- `show_user` (Boolean)
- `tags` (List of String)

