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
	_ datasource.DataSource              = &PanelAnnotationsListDataSource{}
	_ datasource.DataSourceWithConfigure = &PanelAnnotationsListDataSource{}
)

func NewPanelAnnotationsListDataSource() datasource.DataSource {
	return &PanelAnnotationsListDataSource{}
}

// PanelAnnotationsListDataSource defines the data source implementation.
type PanelAnnotationsListDataSource struct{}

// PanelAnnotationsListDataSourceModel describes the data source data model.
type PanelAnnotationsListDataSourceModel struct {
	PanelOptions struct {
		OnlyFromThisDashboard types.Bool   `tfsdk:"only_from_this_dashboard"`
		OnlyInTimeRange       types.Bool   `tfsdk:"only_in_time_range"`
		Tags                  types.List   `tfsdk:"tags"`
		Limit                 types.Int64  `tfsdk:"limit"`
		ShowUser              types.Bool   `tfsdk:"show_user"`
		ShowTime              types.Bool   `tfsdk:"show_time"`
		ShowTags              types.Bool   `tfsdk:"show_tags"`
		NavigateToPanel       types.Bool   `tfsdk:"navigate_to_panel"`
		NavigateBefore        types.String `tfsdk:"navigate_before"`
		NavigateAfter         types.String `tfsdk:"navigate_after"`
	} `tfsdk:"panel_options"`
	Type          types.String `tfsdk:"type"`
	Id            types.Int64  `tfsdk:"id"`
	PluginVersion types.String `tfsdk:"plugin_version"`
	Tags          types.List   `tfsdk:"tags"`
	Targets       []struct {
	} `tfsdk:"targets"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
	Transparent types.Bool   `tfsdk:"transparent"`
	Datasource  *struct {
		Type types.String `tfsdk:"type"`
		Uid  types.String `tfsdk:"uid"`
	} `tfsdk:"datasource"`
	GridPos *struct {
		H      types.Int64 `tfsdk:"h"`
		W      types.Int64 `tfsdk:"w"`
		X      types.Int64 `tfsdk:"x"`
		Y      types.Int64 `tfsdk:"y"`
		Static types.Bool  `tfsdk:"static"`
	} `tfsdk:"grid_pos"`
	Links []struct {
		Title       types.String `tfsdk:"title"`
		Type        types.String `tfsdk:"type"`
		Icon        types.String `tfsdk:"icon"`
		Tooltip     types.String `tfsdk:"tooltip"`
		Url         types.String `tfsdk:"url"`
		Tags        types.List   `tfsdk:"tags"`
		AsDropdown  types.Bool   `tfsdk:"as_dropdown"`
		TargetBlank types.Bool   `tfsdk:"target_blank"`
		IncludeVars types.Bool   `tfsdk:"include_vars"`
		KeepTime    types.Bool   `tfsdk:"keep_time"`
	} `tfsdk:"links"`
	Repeat          types.String  `tfsdk:"repeat"`
	RepeatDirection types.String  `tfsdk:"repeat_direction"`
	RepeatPanelId   types.Int64   `tfsdk:"repeat_panel_id"`
	MaxDataPoints   types.Float64 `tfsdk:"max_data_points"`
	Thresholds      []struct {
	} `tfsdk:"thresholds"`
	TimeRegions []struct {
	} `tfsdk:"time_regions"`
	Transformations []struct {
		Id       types.String `tfsdk:"id"`
		Disabled types.Bool   `tfsdk:"disabled"`
		Filter   *struct {
			Id types.String `tfsdk:"id"`
		} `tfsdk:"filter"`
	} `tfsdk:"transformations"`
	Interval     types.String `tfsdk:"interval"`
	TimeFrom     types.String `tfsdk:"time_from"`
	TimeShift    types.String `tfsdk:"time_shift"`
	LibraryPanel *struct {
		Name types.String `tfsdk:"name"`
		Uid  types.String `tfsdk:"uid"`
	} `tfsdk:"library_panel"`
	Options struct {
	} `tfsdk:"options"`
	FieldConfig struct {
		Defaults struct {
			DisplayName       types.String  `tfsdk:"display_name"`
			DisplayNameFromDS types.String  `tfsdk:"display_name_from_ds"`
			Description       types.String  `tfsdk:"description"`
			Path              types.String  `tfsdk:"path"`
			Writeable         types.Bool    `tfsdk:"writeable"`
			Filterable        types.Bool    `tfsdk:"filterable"`
			Unit              types.String  `tfsdk:"unit"`
			Decimals          types.Float64 `tfsdk:"decimals"`
			Min               types.Float64 `tfsdk:"min"`
			Max               types.Float64 `tfsdk:"max"`
			Thresholds        *struct {
				Mode  types.String `tfsdk:"mode"`
				Steps []struct {
					Value types.Float64 `tfsdk:"value"`
					Color types.String  `tfsdk:"color"`
					State types.String  `tfsdk:"state"`
				} `tfsdk:"steps"`
			} `tfsdk:"thresholds"`
			Color *struct {
				Mode       types.String `tfsdk:"mode"`
				FixedColor types.String `tfsdk:"fixed_color"`
				SeriesBy   types.String `tfsdk:"series_by"`
			} `tfsdk:"color"`
			Links []struct {
			} `tfsdk:"links"`
			NoValue types.String `tfsdk:"no_value"`
			Custom  *struct {
			} `tfsdk:"custom"`
		} `tfsdk:"defaults"`
		Overrides []struct {
			Matcher struct {
				Id types.String `tfsdk:"id"`
			} `tfsdk:"matcher"`
			Properties []struct {
				Id types.String `tfsdk:"id"`
			} `tfsdk:"properties"`
		} `tfsdk:"overrides"`
	} `tfsdk:"field_config"`
	ToJSON types.String `tfsdk:"to_json"`
}

// PanelAnnotationsListDataSourceModelJSON describes the data source data model when exported to json.
type PanelAnnotationsListDataSourceModelJSON struct {
	PanelOptions struct {
		OnlyFromThisDashboard bool     `json:"onlyFromThisDashboard"`
		OnlyInTimeRange       bool     `json:"onlyInTimeRange"`
		Tags                  []string `json:"tags"`
		Limit                 int64    `json:"limit"`
		ShowUser              bool     `json:"showUser"`
		ShowTime              bool     `json:"showTime"`
		ShowTags              bool     `json:"showTags"`
		NavigateToPanel       bool     `json:"navigateToPanel"`
		NavigateBefore        string   `json:"navigateBefore"`
		NavigateAfter         string   `json:"navigateAfter"`
	} `json:"PanelOptions"`
	Type          string   `json:"type"`
	Id            *int64   `json:"id,omitempty"`
	PluginVersion *string  `json:"pluginVersion,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Targets       []struct {
	} `json:"targets,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Transparent bool    `json:"transparent"`
	Datasource  *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	GridPos *struct {
		H      int64 `json:"h"`
		W      int64 `json:"w"`
		X      int64 `json:"x"`
		Y      int64 `json:"y"`
		Static *bool `json:"static,omitempty"`
	} `json:"gridPos,omitempty"`
	Links []struct {
		Title       string   `json:"title"`
		Type        string   `json:"type"`
		Icon        string   `json:"icon"`
		Tooltip     string   `json:"tooltip"`
		Url         string   `json:"url"`
		Tags        []string `json:"tags"`
		AsDropdown  bool     `json:"asDropdown"`
		TargetBlank bool     `json:"targetBlank"`
		IncludeVars bool     `json:"includeVars"`
		KeepTime    bool     `json:"keepTime"`
	} `json:"links,omitempty"`
	Repeat          *string  `json:"repeat,omitempty"`
	RepeatDirection string   `json:"repeatDirection"`
	RepeatPanelId   *int64   `json:"repeatPanelId,omitempty"`
	MaxDataPoints   *float64 `json:"maxDataPoints,omitempty"`
	Thresholds      []struct {
	} `json:"thresholds,omitempty"`
	TimeRegions []struct {
	} `json:"timeRegions,omitempty"`
	Transformations []struct {
		Id       string `json:"id"`
		Disabled *bool  `json:"disabled,omitempty"`
		Filter   *struct {
			Id string `json:"id"`
		} `json:"filter,omitempty"`
	} `json:"transformations"`
	Interval     *string `json:"interval,omitempty"`
	TimeFrom     *string `json:"timeFrom,omitempty"`
	TimeShift    *string `json:"timeShift,omitempty"`
	LibraryPanel *struct {
		Name string `json:"name"`
		Uid  string `json:"uid"`
	} `json:"libraryPanel,omitempty"`
	Options struct {
	} `json:"options"`
	FieldConfig struct {
		Defaults struct {
			DisplayName       *string  `json:"displayName,omitempty"`
			DisplayNameFromDS *string  `json:"displayNameFromDS,omitempty"`
			Description       *string  `json:"description,omitempty"`
			Path              *string  `json:"path,omitempty"`
			Writeable         *bool    `json:"writeable,omitempty"`
			Filterable        *bool    `json:"filterable,omitempty"`
			Unit              *string  `json:"unit,omitempty"`
			Decimals          *float64 `json:"decimals,omitempty"`
			Min               *float64 `json:"min,omitempty"`
			Max               *float64 `json:"max,omitempty"`
			Thresholds        *struct {
				Mode  string `json:"mode"`
				Steps []struct {
					Value *float64 `json:"value,omitempty"`
					Color string   `json:"color"`
					State *string  `json:"state,omitempty"`
				} `json:"steps"`
			} `json:"thresholds,omitempty"`
			Color *struct {
				Mode       string  `json:"mode"`
				FixedColor *string `json:"fixedColor,omitempty"`
				SeriesBy   *string `json:"seriesBy,omitempty"`
			} `json:"color,omitempty"`
			Links []struct {
			} `json:"links,omitempty"`
			NoValue *string `json:"noValue,omitempty"`
			Custom  *struct {
			} `json:"custom,omitempty"`
		} `json:"defaults"`
		Overrides []struct {
			Matcher struct {
				Id string `json:"id"`
			} `json:"matcher"`
			Properties []struct {
				Id string `json:"id"`
			} `json:"properties"`
		} `json:"overrides"`
	} `json:"fieldConfig"`
}

func (d *PanelAnnotationsListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_panel_annotations_list"
}

func (d *PanelAnnotationsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
					"only_from_this_dashboard": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"only_in_time_range": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
						ElementType:         types.StringType,
					},
					"limit": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"show_user": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"show_time": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"show_tags": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"navigate_to_panel": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"navigate_before": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"navigate_after": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            true,
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
			"max_data_points": schema.Float64Attribute{
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
							"decimals": schema.Float64Attribute{
								MarkdownDescription: `Significant digits (for display)`,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"min": schema.Float64Attribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
							"max": schema.Float64Attribute{
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
												"value": schema.Float64Attribute{
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

func (d *PanelAnnotationsListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *PanelAnnotationsListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PanelAnnotationsListDataSourceModel

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

func (d *PanelAnnotationsListDataSource) applyDefaults(data *PanelAnnotationsListDataSourceModel) {
	if data.PanelOptions.OnlyFromThisDashboard.IsNull() {
		data.PanelOptions.OnlyFromThisDashboard = types.BoolValue(false)
	}
	if data.PanelOptions.OnlyInTimeRange.IsNull() {
		data.PanelOptions.OnlyInTimeRange = types.BoolValue(false)
	}
	if data.PanelOptions.Limit.IsNull() {
		data.PanelOptions.Limit = types.Int64Value(10)
	}
	if data.PanelOptions.ShowUser.IsNull() {
		data.PanelOptions.ShowUser = types.BoolValue(true)
	}
	if data.PanelOptions.ShowTime.IsNull() {
		data.PanelOptions.ShowTime = types.BoolValue(true)
	}
	if data.PanelOptions.ShowTags.IsNull() {
		data.PanelOptions.ShowTags = types.BoolValue(true)
	}
	if data.PanelOptions.NavigateToPanel.IsNull() {
		data.PanelOptions.NavigateToPanel = types.BoolValue(true)
	}
	if data.PanelOptions.NavigateBefore.IsNull() {
		data.PanelOptions.NavigateBefore = types.StringValue(`10m`)
	}
	if data.PanelOptions.NavigateAfter.IsNull() {
		data.PanelOptions.NavigateAfter = types.StringValue(`10m`)
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

func (d PanelAnnotationsListDataSourceModel) MarshalJSON() ([]byte, error) {
	attr_type := d.Type.ValueString()
	attr_id := d.Id.ValueInt64()
	attr_pluginversion := d.PluginVersion.ValueString()
	attr_title := d.Title.ValueString()
	attr_description := d.Description.ValueString()
	attr_transparent := d.Transparent.ValueBool()
	attr_repeat := d.Repeat.ValueString()
	attr_repeatdirection := d.RepeatDirection.ValueString()
	attr_repeatpanelid := d.RepeatPanelId.ValueInt64()
	attr_maxdatapoints := d.MaxDataPoints.ValueFloat64()
	attr_interval := d.Interval.ValueString()
	attr_timefrom := d.TimeFrom.ValueString()
	attr_timeshift := d.TimeShift.ValueString()

	model := &PanelAnnotationsListDataSourceModelJSON{
		Type:            attr_type,
		Id:              &attr_id,
		PluginVersion:   &attr_pluginversion,
		Title:           &attr_title,
		Description:     &attr_description,
		Transparent:     attr_transparent,
		Repeat:          &attr_repeat,
		RepeatDirection: attr_repeatdirection,
		RepeatPanelId:   &attr_repeatpanelid,
		MaxDataPoints:   &attr_maxdatapoints,
		Interval:        &attr_interval,
		TimeFrom:        &attr_timefrom,
		TimeShift:       &attr_timeshift,
	}
	return json.Marshal(model)
}
