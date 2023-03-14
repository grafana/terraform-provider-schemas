// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformDataSourceJenny
//     ComposableLatestMajorsOrXJenny
//
// Run 'go generate ./' from repository root to regenerate.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &PanelHistogramDataSource{}
	_ datasource.DataSourceWithConfigure = &PanelHistogramDataSource{}
)

func NewPanelHistogramDataSource() datasource.DataSource {
	return &PanelHistogramDataSource{}
}

// PanelHistogramDataSource defines the data source implementation.
type PanelHistogramDataSource struct{}

// PanelHistogramDataSourceModel describes the data source data model.
type PanelHistogramDataSourceModel struct {
	PanelOptions struct {
		BucketSize   types.Int64 `tfsdk:"bucket_size" json:"bucketSize"`
		BucketOffset types.Int64 `tfsdk:"bucket_offset" json:"bucketOffset"`
		Legend       struct {
			DisplayMode types.String `tfsdk:"display_mode" json:"displayMode"`
			Placement   types.String `tfsdk:"placement" json:"placement"`
			ShowLegend  types.Bool   `tfsdk:"show_legend" json:"showLegend"`
			AsTable     types.Bool   `tfsdk:"as_table" json:"asTable"`
			IsVisible   types.Bool   `tfsdk:"is_visible" json:"isVisible"`
			SortBy      types.String `tfsdk:"sort_by" json:"sortBy"`
			SortDesc    types.Bool   `tfsdk:"sort_desc" json:"sortDesc"`
			Width       types.Number `tfsdk:"width" json:"width"`
			Calcs       types.List   `tfsdk:"calcs" json:"calcs"`
		} `tfsdk:"legend" json:"legend"`
		Tooltip struct {
			Mode types.String `tfsdk:"mode" json:"mode"`
			Sort types.String `tfsdk:"sort" json:"sort"`
		} `tfsdk:"tooltip" json:"tooltip"`
		Combine types.Bool `tfsdk:"combine" json:"combine"`
	} `tfsdk:"panel_options" json:"PanelOptions"`
	PanelFieldConfig struct {
		LineWidth         types.Int64  `tfsdk:"line_width" json:"lineWidth"`
		FillOpacity       types.Int64  `tfsdk:"fill_opacity" json:"fillOpacity"`
		AxisPlacement     types.String `tfsdk:"axis_placement" json:"axisPlacement"`
		AxisColorMode     types.String `tfsdk:"axis_color_mode" json:"axisColorMode"`
		AxisLabel         types.String `tfsdk:"axis_label" json:"axisLabel"`
		AxisWidth         types.Number `tfsdk:"axis_width" json:"axisWidth"`
		AxisSoftMin       types.Number `tfsdk:"axis_soft_min" json:"axisSoftMin"`
		AxisSoftMax       types.Number `tfsdk:"axis_soft_max" json:"axisSoftMax"`
		AxisGridShow      types.Bool   `tfsdk:"axis_grid_show" json:"axisGridShow"`
		ScaleDistribution *struct {
			Type            types.String `tfsdk:"type" json:"type"`
			Log             types.Number `tfsdk:"log" json:"log"`
			LinearThreshold types.Number `tfsdk:"linear_threshold" json:"linearThreshold"`
		} `tfsdk:"scale_distribution" json:"scaleDistribution"`
		HideFrom *struct {
			Tooltip types.Bool `tfsdk:"tooltip" json:"tooltip"`
			Legend  types.Bool `tfsdk:"legend" json:"legend"`
			Viz     types.Bool `tfsdk:"viz" json:"viz"`
		} `tfsdk:"hide_from" json:"hideFrom"`
		GradientMode     types.String `tfsdk:"gradient_mode" json:"gradientMode"`
		AxisCenteredZero types.Bool   `tfsdk:"axis_centered_zero" json:"axisCenteredZero"`
	} `tfsdk:"panel_field_config" json:"PanelFieldConfig"`
	Type          types.String `tfsdk:"type" json:"type"`
	Id            types.Int64  `tfsdk:"id" json:"id"`
	PluginVersion types.String `tfsdk:"plugin_version" json:"pluginVersion"`
	Tags          types.List   `tfsdk:"tags" json:"tags"`
	Targets       []struct {
	} `tfsdk:"targets" json:"targets"`
	Title       types.String `tfsdk:"title" json:"title"`
	Description types.String `tfsdk:"description" json:"description"`
	Transparent types.Bool   `tfsdk:"transparent" json:"transparent"`
	Datasource  *struct {
		Type types.String `tfsdk:"type" json:"type"`
		Uid  types.String `tfsdk:"uid" json:"uid"`
	} `tfsdk:"datasource" json:"datasource"`
	GridPos *struct {
		H      types.Int64 `tfsdk:"h" json:"h"`
		W      types.Int64 `tfsdk:"w" json:"w"`
		X      types.Int64 `tfsdk:"x" json:"x"`
		Y      types.Int64 `tfsdk:"y" json:"y"`
		Static types.Bool  `tfsdk:"static" json:"static"`
	} `tfsdk:"grid_pos" json:"gridPos"`
	Links []struct {
		Title       types.String `tfsdk:"title" json:"title"`
		Type        types.String `tfsdk:"type" json:"type"`
		Icon        types.String `tfsdk:"icon" json:"icon"`
		Tooltip     types.String `tfsdk:"tooltip" json:"tooltip"`
		Url         types.String `tfsdk:"url" json:"url"`
		Tags        types.List   `tfsdk:"tags" json:"tags"`
		AsDropdown  types.Bool   `tfsdk:"as_dropdown" json:"asDropdown"`
		TargetBlank types.Bool   `tfsdk:"target_blank" json:"targetBlank"`
		IncludeVars types.Bool   `tfsdk:"include_vars" json:"includeVars"`
		KeepTime    types.Bool   `tfsdk:"keep_time" json:"keepTime"`
	} `tfsdk:"links" json:"links"`
	Repeat          types.String `tfsdk:"repeat" json:"repeat"`
	RepeatDirection types.String `tfsdk:"repeat_direction" json:"repeatDirection"`
	RepeatPanelId   types.Int64  `tfsdk:"repeat_panel_id" json:"repeatPanelId"`
	MaxDataPoints   types.Number `tfsdk:"max_data_points" json:"maxDataPoints"`
	Thresholds      []struct {
	} `tfsdk:"thresholds" json:"thresholds"`
	TimeRegions []struct {
	} `tfsdk:"time_regions" json:"timeRegions"`
	Transformations []struct {
		Id       types.String `tfsdk:"id" json:"id"`
		Disabled types.Bool   `tfsdk:"disabled" json:"disabled"`
		Filter   *struct {
			Id types.String `tfsdk:"id" json:"id"`
		} `tfsdk:"filter" json:"filter"`
	} `tfsdk:"transformations" json:"transformations"`
	Interval     types.String `tfsdk:"interval" json:"interval"`
	TimeFrom     types.String `tfsdk:"time_from" json:"timeFrom"`
	TimeShift    types.String `tfsdk:"time_shift" json:"timeShift"`
	LibraryPanel *struct {
		Name types.String `tfsdk:"name" json:"name"`
		Uid  types.String `tfsdk:"uid" json:"uid"`
	} `tfsdk:"library_panel" json:"libraryPanel"`
	Options struct {
	} `tfsdk:"options" json:"options"`
	FieldConfig struct {
		Defaults struct {
			DisplayName       types.String `tfsdk:"display_name" json:"displayName"`
			DisplayNameFromDS types.String `tfsdk:"display_name_from_ds" json:"displayNameFromDS"`
			Description       types.String `tfsdk:"description" json:"description"`
			Path              types.String `tfsdk:"path" json:"path"`
			Writeable         types.Bool   `tfsdk:"writeable" json:"writeable"`
			Filterable        types.Bool   `tfsdk:"filterable" json:"filterable"`
			Unit              types.String `tfsdk:"unit" json:"unit"`
			Decimals          types.Number `tfsdk:"decimals" json:"decimals"`
			Min               types.Number `tfsdk:"min" json:"min"`
			Max               types.Number `tfsdk:"max" json:"max"`
			Thresholds        *struct {
				Mode  types.String `tfsdk:"mode" json:"mode"`
				Steps []struct {
					Value types.Number `tfsdk:"value" json:"value"`
					Color types.String `tfsdk:"color" json:"color"`
					State types.String `tfsdk:"state" json:"state"`
				} `tfsdk:"steps" json:"steps"`
			} `tfsdk:"thresholds" json:"thresholds"`
			Color *struct {
				Mode       types.String `tfsdk:"mode" json:"mode"`
				FixedColor types.String `tfsdk:"fixed_color" json:"fixedColor"`
				SeriesBy   types.String `tfsdk:"series_by" json:"seriesBy"`
			} `tfsdk:"color" json:"color"`
			Links []struct {
			} `tfsdk:"links" json:"links"`
			NoValue types.String `tfsdk:"no_value" json:"noValue"`
			Custom  *struct {
			} `tfsdk:"custom" json:"custom"`
		} `tfsdk:"defaults" json:"defaults"`
		Overrides []struct {
			Matcher struct {
				Id types.String `tfsdk:"id" json:"id"`
			} `tfsdk:"matcher" json:"matcher"`
			Properties []struct {
				Id types.String `tfsdk:"id" json:"id"`
			} `tfsdk:"properties" json:"properties"`
		} `tfsdk:"overrides" json:"overrides"`
	} `tfsdk:"field_config" json:"fieldConfig"`
	ToJSON types.String `tfsdk:"to_json"`
}

func (d *PanelHistogramDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_panel_histogram"
}

func (d *PanelHistogramDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",
		Attributes: map[string]schema.Attribute{
			"panel_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"bucket_size": schema.Int64Attribute{
						MarkdownDescription: `Size of each bucket`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"bucket_offset": schema.Int64Attribute{
						MarkdownDescription: `Offset buckets by this amount`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"legend": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"display_mode": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"placement": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"show_legend": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"as_table": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"is_visible": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"sort_by": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"sort_desc": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"width": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"calcs": schema.ListAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
								ElementType:         types.StringType,
							},
						},
					},
					"tooltip": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"sort": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
						},
					},
					"combine": schema.BoolAttribute{
						MarkdownDescription: `Combines multiple series into a single histogram`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"panel_field_config": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"line_width": schema.Int64Attribute{
						MarkdownDescription: `Controls line width of the bars.`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"fill_opacity": schema.Int64Attribute{
						MarkdownDescription: `Controls the fill opacity of the bars.`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"axis_placement": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_color_mode": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_label": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_width": schema.NumberAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_soft_min": schema.NumberAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_soft_max": schema.NumberAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"axis_grid_show": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"scale_distribution": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"log": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"linear_threshold": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
						},
					},
					"hide_from": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"tooltip": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"legend": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"viz": schema.BoolAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
						},
					},
					"gradient_mode": schema.StringAttribute{
						MarkdownDescription: `Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
Gradient appearance is influenced by the Fill opacity setting.`,
						Computed: true,
						Optional: true,
						Required: false,
					},
					"axis_centered_zero": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The panel plugin type id. May not be empty.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: `TODO docs`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"plugin_version": schema.StringAttribute{
				MarkdownDescription: `FIXME this almost certainly has to be changed in favor of scuemata versions`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `TODO docs`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				ElementType:         types.StringType,
			},
			"targets": schema.ListNestedAttribute{
				MarkdownDescription: `TODO docs`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: `Panel title.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Description.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"transparent": schema.BoolAttribute{
				MarkdownDescription: `Whether to display the panel without a background.`,
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"datasource": schema.SingleNestedAttribute{
				MarkdownDescription: `The datasource used in all targets.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"uid": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"grid_pos": schema.SingleNestedAttribute{
				MarkdownDescription: `Grid position.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"h": schema.Int64Attribute{
						MarkdownDescription: `Panel`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"w": schema.Int64Attribute{
						MarkdownDescription: `Panel`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"x": schema.Int64Attribute{
						MarkdownDescription: `Panel x`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"y": schema.Int64Attribute{
						MarkdownDescription: `Panel y`,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"static": schema.BoolAttribute{
						MarkdownDescription: `true if fixed`,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"links": schema.ListNestedAttribute{
				MarkdownDescription: `Panel links.
TODO fill this out - seems there are a couple variants?`,
				Computed: false,
				Optional: true,
				Required: false,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"title": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"icon": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"tooltip": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            false,
							Required:            true,
							ElementType:         types.StringType,
						},
						"as_dropdown": schema.BoolAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							Optional:            true,
							Required:            false,
						},
						"target_blank": schema.BoolAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							Optional:            true,
							Required:            false,
						},
						"include_vars": schema.BoolAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							Optional:            true,
							Required:            false,
						},
						"keep_time": schema.BoolAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							Optional:            true,
							Required:            false,
						},
					},
				},
			},
			"repeat": schema.StringAttribute{
				MarkdownDescription: `Name of template variable to repeat for.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"repeat_direction": schema.StringAttribute{
				MarkdownDescription: `Direction to repeat in if 'repeat' is set.
"h" for horizontal, "v" for vertical.
TODO this is probably optional`,
				Computed: true,
				Optional: true,
				Required: false,
			},
			"repeat_panel_id": schema.Int64Attribute{
				MarkdownDescription: `Id of the repeating panel.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"max_data_points": schema.NumberAttribute{
				MarkdownDescription: `TODO docs`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"thresholds": schema.ListNestedAttribute{
				MarkdownDescription: `TODO docs - seems to be an old field from old dashboard alerts?`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"time_regions": schema.ListNestedAttribute{
				MarkdownDescription: `TODO docs`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"transformations": schema.ListNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: `Unique identifier of transformer`,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"disabled": schema.BoolAttribute{
							MarkdownDescription: `Disabled transformations are skipped`,
							Computed:            false,
							Optional:            true,
							Required:            false,
						},
						"filter": schema.SingleNestedAttribute{
							MarkdownDescription: `Optional frame matcher.  When missing it will be applied to all results`,
							Computed:            false,
							Optional:            true,
							Required:            false,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									MarkdownDescription: ``,
									Computed:            true,
									Optional:            true,
									Required:            false,
								},
							},
						},
					},
				},
			},
			"interval": schema.StringAttribute{
				MarkdownDescription: `TODO docs
TODO tighter constraint`,
				Computed: false,
				Optional: true,
				Required: false,
			},
			"time_from": schema.StringAttribute{
				MarkdownDescription: `TODO docs
TODO tighter constraint`,
				Computed: false,
				Optional: true,
				Required: false,
			},
			"time_shift": schema.StringAttribute{
				MarkdownDescription: `TODO docs
TODO tighter constraint`,
				Computed: false,
				Optional: true,
				Required: false,
			},
			"library_panel": schema.SingleNestedAttribute{
				MarkdownDescription: `Dynamically load the panel`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"uid": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
				},
			},
			"options": schema.SingleNestedAttribute{
				MarkdownDescription: `options is specified by the PanelOptions field in panel
plugin schemas.`,
				Computed: false,
				Optional: false,
				Required: true,
			},
			"field_config": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"defaults": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"display_name": schema.StringAttribute{
								MarkdownDescription: `The display value for this field.  This supports template variables blank is auto`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"display_name_from_ds": schema.StringAttribute{
								MarkdownDescription: `This can be used by data sources that return and explicit naming structure for values and labels
When this property is configured, this value is used rather than the default naming strategy.`,
								Computed: false,
								Optional: true,
								Required: false,
							},
							"description": schema.StringAttribute{
								MarkdownDescription: `Human readable field metadata`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"path": schema.StringAttribute{
								MarkdownDescription: `An explicit path to the field in the datasource.  When the frame meta includes a path,
This will default to ${frame.meta.path}/${field.name}

When defined, this value can be used as an identifier within the datasource scope, and
may be used to update the results`,
								Computed: false,
								Optional: true,
								Required: false,
							},
							"writeable": schema.BoolAttribute{
								MarkdownDescription: `True if data source can write a value to the path.  Auth/authz are supported separately`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"filterable": schema.BoolAttribute{
								MarkdownDescription: `True if data source field supports ad-hoc filters`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"unit": schema.StringAttribute{
								MarkdownDescription: `Numeric Options`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"decimals": schema.NumberAttribute{
								MarkdownDescription: `Significant digits (for display)`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"min": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"max": schema.NumberAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"thresholds": schema.SingleNestedAttribute{
								MarkdownDescription: `Map numeric values to states`,
								Computed:            false,
								Optional:            true,
								Required:            false,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										MarkdownDescription: ``,
										Computed:            false,
										Optional:            false,
										Required:            true,
									},
									"steps": schema.ListNestedAttribute{
										MarkdownDescription: `Must be sorted by 'value', first value is always -Infinity`,
										Computed:            false,
										Optional:            false,
										Required:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"value": schema.NumberAttribute{
													MarkdownDescription: `TODO docs
FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON`,
													Computed: false,
													Optional: true,
													Required: false,
												},
												"color": schema.StringAttribute{
													MarkdownDescription: `TODO docs`,
													Computed:            false,
													Optional:            false,
													Required:            true,
												},
												"state": schema.StringAttribute{
													MarkdownDescription: `TODO docs
TODO are the values here enumerable into a disjunction?
Some seem to be listed in typescript comment`,
													Computed: false,
													Optional: true,
													Required: false,
												},
											},
										},
									},
								},
							},
							"color": schema.SingleNestedAttribute{
								MarkdownDescription: `Map values to a display color`,
								Computed:            false,
								Optional:            true,
								Required:            false,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										MarkdownDescription: `The main color scheme mode`,
										Computed:            false,
										Optional:            false,
										Required:            true,
									},
									"fixed_color": schema.StringAttribute{
										MarkdownDescription: `Stores the fixed color value if mode is fixed`,
										Computed:            false,
										Optional:            true,
										Required:            false,
									},
									"series_by": schema.StringAttribute{
										MarkdownDescription: `Some visualizations need to know how to assign a series color from by value color schemes`,
										Computed:            false,
										Optional:            true,
										Required:            false,
									},
								},
							},
							"links": schema.ListNestedAttribute{
								MarkdownDescription: `The behavior when clicking on a result`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"no_value": schema.StringAttribute{
								MarkdownDescription: `Alternative to empty string`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"custom": schema.SingleNestedAttribute{
								MarkdownDescription: `custom is specified by the PanelFieldConfig field
in panel plugin schemas.`,
								Computed: false,
								Optional: true,
								Required: false,
							},
						},
					},
					"overrides": schema.ListNestedAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"matcher": schema.SingleNestedAttribute{
									MarkdownDescription: ``,
									Computed:            false,
									Optional:            false,
									Required:            true,
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											MarkdownDescription: ``,
											Computed:            true,
											Optional:            true,
											Required:            false,
										},
									},
								},
								"properties": schema.ListNestedAttribute{
									MarkdownDescription: ``,
									Computed:            false,
									Optional:            false,
									Required:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												MarkdownDescription: ``,
												Computed:            true,
												Optional:            true,
												Required:            false,
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *PanelHistogramDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *PanelHistogramDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PanelHistogramDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	d.applyDefaults(&data)
	JSONConfig, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("JSON marshalling error", err.Error())
		return
	}

	// Not sure about that
	data.ToJSON = types.StringValue(string(JSONConfig))

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *PanelHistogramDataSource) applyDefaults(data *PanelHistogramDataSourceModel) {
	if data.PanelOptions.BucketOffset.IsNull() {
		data.PanelOptions.BucketOffset = types.Int64Value(0)
	}
	if data.PanelFieldConfig.LineWidth.IsNull() {
		data.PanelFieldConfig.LineWidth = types.Int64Value(1)
	}
	if data.PanelFieldConfig.FillOpacity.IsNull() {
		data.PanelFieldConfig.FillOpacity = types.Int64Value(80)
	}
	if data.PanelFieldConfig.GradientMode.IsNull() {
		data.PanelFieldConfig.GradientMode = types.StringValue(`none`)
	}
	if data.Transparent.IsNull() {
		data.Transparent = types.BoolValue(false)
	}
	if data.GridPos != nil && data.GridPos.H.IsNull() {
		data.GridPos.H = types.Int64Value(9)
	}
	if data.GridPos != nil && data.GridPos.W.IsNull() {
		data.GridPos.W = types.Int64Value(12)
	}
	if data.GridPos != nil && data.GridPos.X.IsNull() {
		data.GridPos.X = types.Int64Value(0)
	}
	if data.GridPos != nil && data.GridPos.Y.IsNull() {
		data.GridPos.Y = types.Int64Value(0)
	}
	if data.RepeatDirection.IsNull() {
		data.RepeatDirection = types.StringValue(`h`)
	}
}
