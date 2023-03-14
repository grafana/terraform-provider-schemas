---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "schemas_core_preferences Data Source - terraform-provider-schemas"
subcategory: ""
description: |-
  TODO description
---

# schemas_core_preferences (Data Source)

TODO description



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `home_dashboard_uid` (String) UID for the home dashboard
- `language` (String) Selected language (beta)
- `query_history` (Attributes) Explore query history preferences (see [below for nested schema](#nestedatt--query_history))
- `theme` (String) light, dark, empty is default
- `timezone` (String) The timezone selection
TODO: this should use the timezone defined in common
- `week_start` (String) day of the week (sunday, monday, etc)

### Read-Only

- `to_json` (String) This datasource rendered as JSON

<a id="nestedatt--query_history"></a>
### Nested Schema for `query_history`

Optional:

- `home_tab` (String) one of: '' | 'query' | 'starred';

