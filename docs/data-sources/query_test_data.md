---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "schemas_query_test_data Data Source - terraform-provider-schemas"
subcategory: ""
description: |-
  TODO description
---

# schemas_query_test_data (Data Source)

TODO description



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ref_id` (String) A - Z

### Optional

- `alias` (String)
- `channel` (String)
- `csv_content` (String)
- `csv_file_name` (String)
- `csv_wave` (Attributes List) (see [below for nested schema](#nestedatt--csv_wave))
- `error_type` (String)
- `hide` (Boolean) true if query is disabled (ie should not be returned to the dashboard)
- `key` (String) Unique, guid like, string used in explore mode
- `labels` (String)
- `level_column` (Boolean)
- `lines` (Number)
- `nodes` (Attributes) (see [below for nested schema](#nestedatt--nodes))
- `pulse_wave` (Attributes) (see [below for nested schema](#nestedatt--pulse_wave))
- `query_type` (String) Specify the query flavor
TODO make this required and give it a default
- `raw_frame_content` (String)
- `scenario_id` (String)
- `series_count` (Number)
- `sim` (Attributes) (see [below for nested schema](#nestedatt--sim))
- `span_count` (Number)
- `stream` (Attributes) (see [below for nested schema](#nestedatt--stream))
- `string_input` (String)
- `usa` (Attributes) (see [below for nested schema](#nestedatt--usa))

### Read-Only

- `to_json` (String) This datasource rendered as JSON

<a id="nestedatt--csv_wave"></a>
### Nested Schema for `csv_wave`

Optional:

- `labels` (String)
- `name` (String)
- `time_step` (Number)
- `values_csv` (String)


<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Optional:

- `count` (Number)
- `type` (String)


<a id="nestedatt--pulse_wave"></a>
### Nested Schema for `pulse_wave`

Optional:

- `off_count` (Number)
- `off_value` (Number)
- `on_count` (Number)
- `on_value` (Number)
- `time_step` (Number)


<a id="nestedatt--sim"></a>
### Nested Schema for `sim`

Required:

- `key` (Attributes) (see [below for nested schema](#nestedatt--sim--key))

Optional:

- `config` (Attributes) (see [below for nested schema](#nestedatt--sim--config))
- `last` (Boolean)
- `stream` (Boolean)

<a id="nestedatt--sim--key"></a>
### Nested Schema for `sim.key`

Required:

- `tick` (Number)
- `type` (String)

Optional:

- `uid` (String)


<a id="nestedatt--sim--config"></a>
### Nested Schema for `sim.config`



<a id="nestedatt--stream"></a>
### Nested Schema for `stream`

Required:

- `noise` (Number)
- `speed` (Number)
- `spread` (Number)
- `type` (String)

Optional:

- `bands` (Number)
- `url` (String)


<a id="nestedatt--usa"></a>
### Nested Schema for `usa`

Optional:

- `fields` (List of String)
- `mode` (String)
- `period` (String)
- `states` (List of String)

