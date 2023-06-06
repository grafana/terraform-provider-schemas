// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformDataSourceJenny
//     LatestJenny
//
// Run 'go generate ./' from repository root to regenerate.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure that the imports are used to avoid compiler errors.
var _ attr.Value
var _ diag.Diagnostic

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &CorePlaylistDataSource{}
	_ datasource.DataSourceWithConfigure = &CorePlaylistDataSource{}
)

func NewCorePlaylistDataSource() datasource.DataSource {
	return &CorePlaylistDataSource{}
}

// CorePlaylistDataSource defines the data source implementation.
type CorePlaylistDataSource struct{}

type CorePlaylistDataSourceModel_Items struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
	Title types.String `tfsdk:"title"`
}

func (m CorePlaylistDataSourceModel_Items) MarshalJSON() ([]byte, error) {
	type jsonCorePlaylistDataSourceModel_Items struct {
		Type  string  `json:"type"`
		Value string  `json:"value"`
		Title *string `json:"title,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_type := m.Type.ValueString()
	attr_value := m.Value.ValueString()
	attr_title := m.Title.ValueString()

	model := &jsonCorePlaylistDataSourceModel_Items{
		Type:  attr_type,
		Value: attr_value,
		Title: &attr_title,
	}
	return json.Marshal(model)
}

func (m CorePlaylistDataSourceModel_Items) ApplyDefaults() CorePlaylistDataSourceModel_Items {

	return m
}

type CorePlaylistDataSourceModel struct {
	ToJSON   types.String                        `tfsdk:"to_json"`
	Uid      types.String                        `tfsdk:"uid"`
	Name     types.String                        `tfsdk:"name"`
	Interval types.String                        `tfsdk:"interval"`
	Items    []CorePlaylistDataSourceModel_Items `tfsdk:"items"`
}

func (m CorePlaylistDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonCorePlaylistDataSourceModel struct {
		Uid      string        `json:"uid"`
		Name     string        `json:"name"`
		Interval string        `json:"interval"`
		Items    []interface{} `json:"items,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_uid := m.Uid.ValueString()
	attr_name := m.Name.ValueString()
	attr_interval := m.Interval.ValueString()
	attr_items := []interface{}{}
	for _, v := range m.Items {
		attr_items = append(attr_items, v)
	}

	model := &jsonCorePlaylistDataSourceModel{
		Uid:      attr_uid,
		Name:     attr_name,
		Interval: attr_interval,
		Items:    attr_items,
	}
	return json.Marshal(model)
}

func (m CorePlaylistDataSourceModel) ApplyDefaults() CorePlaylistDataSourceModel {
	if m.Interval.IsNull() {
		m.Interval = types.StringValue(`5m`)
	}
	return m
}

func (d *CorePlaylistDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_playlist"
}

func (d *CorePlaylistDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"uid": schema.StringAttribute{
				MarkdownDescription: `Unique playlist identifier. Generated on creation, either by the
creator of the playlist of by the application.`,
				Computed: false,
				Optional: false,
				Required: true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the playlist.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"interval": schema.StringAttribute{
				MarkdownDescription: `Interval sets the time between switching views in a playlist.
FIXME: Is this based on a standardized format or what options are available? Can datemath be used?. Defaults to "5m".`,
				Computed: true,
				Optional: true,
				Required: false,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: `The ordered list of items that the playlist will iterate over.
FIXME! This should not be optional, but changing it makes the godegen awkward`,
				Computed: false,
				Optional: true,
				Required: false,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: `Type of the item.`,
							Computed:            false,
							Optional:            false,
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `Value depends on type and describes the playlist item.

 - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
 is not portable as the numerical identifier is non-deterministic between different instances.
 Will be replaced by dashboard_by_uid in the future. (deprecated)
 - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
 dashboards behind the tag will be added to the playlist.
 - dashboard_by_uid: The value is the dashboard UID`,
							Computed: false,
							Optional: false,
							Required: true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: `Title is an unused property -- it will be removed in the future`,
							Computed:            false,
							Optional:            true,
							Required:            false,
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

func (d *CorePlaylistDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CorePlaylistDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CorePlaylistDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

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