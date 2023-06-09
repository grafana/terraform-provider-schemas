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
	_ datasource.DataSource              = &QueryGrafanaPyroscopeDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryGrafanaPyroscopeDataSource{}
)

func NewQueryGrafanaPyroscopeDataSource() datasource.DataSource {
	return &QueryGrafanaPyroscopeDataSource{}
}

// QueryGrafanaPyroscopeDataSource defines the data source implementation.
type QueryGrafanaPyroscopeDataSource struct{}

type QueryGrafanaPyroscopeDataSourceModel struct {
	RenderedJSON  types.String `tfsdk:"rendered_json"`
	LabelSelector types.String `tfsdk:"label_selector"`
	ProfileTypeId types.String `tfsdk:"profile_type_id"`
	GroupBy       types.List   `tfsdk:"group_by"`
	MaxNodes      types.Int64  `tfsdk:"max_nodes"`
	RefId         types.String `tfsdk:"ref_id"`
	Hide          types.Bool   `tfsdk:"hide"`
	QueryType     types.String `tfsdk:"query_type"`
}

func (m QueryGrafanaPyroscopeDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryGrafanaPyroscopeDataSourceModel struct {
		LabelSelector string   `json:"labelSelector"`
		ProfileTypeId string   `json:"profileTypeId"`
		GroupBy       []string `json:"groupBy,omitempty"`
		MaxNodes      *int64   `json:"maxNodes,omitempty"`
		RefId         string   `json:"refId"`
		Hide          *bool    `json:"hide,omitempty"`
		QueryType     *string  `json:"queryType,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_labelselector := m.LabelSelector.ValueString()
	attr_profiletypeid := m.ProfileTypeId.ValueString()
	attr_groupby := []string{}
	for _, v := range m.GroupBy.Elements() {
		attr_groupby = append(attr_groupby, v.(types.String).ValueString())
	}
	attr_maxnodes := m.MaxNodes.ValueInt64()
	attr_refid := m.RefId.ValueString()
	attr_hide := m.Hide.ValueBool()
	attr_querytype := m.QueryType.ValueString()

	model := &jsonQueryGrafanaPyroscopeDataSourceModel{
		LabelSelector: attr_labelselector,
		ProfileTypeId: attr_profiletypeid,
		GroupBy:       attr_groupby,
		MaxNodes:      &attr_maxnodes,
		RefId:         attr_refid,
		Hide:          &attr_hide,
		QueryType:     &attr_querytype,
	}
	return json.Marshal(model)
}

func (m QueryGrafanaPyroscopeDataSourceModel) ApplyDefaults() QueryGrafanaPyroscopeDataSourceModel {
	if m.LabelSelector.IsNull() {
		m.LabelSelector = types.StringValue(`{}`)
	}
	if len(m.GroupBy.Elements()) == 0 {
		m.GroupBy, _ = types.ListValue(types.StringType, []attr.Value{})
	}
	return m
}

func (d *QueryGrafanaPyroscopeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_grafana_pyroscope"
}

func (d *QueryGrafanaPyroscopeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"label_selector": schema.StringAttribute{
				MarkdownDescription: `Specifies the query label selectors. Defaults to "{}".`,
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"profile_type_id": schema.StringAttribute{
				MarkdownDescription: `Specifies the type of profile to query.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"group_by": schema.ListAttribute{
				MarkdownDescription: `Allows to group the results.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				ElementType:         types.StringType,
			},
			"max_nodes": schema.Int64Attribute{
				MarkdownDescription: `Sets the maximum number of nodes in the flamegraph.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"ref_id": schema.StringAttribute{
				MarkdownDescription: `A unique identifier for the query within the list of targets.
In server side expressions, the refId is used as a variable name to identify results.
By default, the UI will assign A->Z; however setting meaningful names may be useful.`,
				Computed: false,
				Optional: false,
				Required: true,
			},
			"hide": schema.BoolAttribute{
				MarkdownDescription: `true if query is disabled (ie should not be returned to the dashboard)
Note this does not always imply that the query should not be executed since
the results from a hidden query may be used as the input to other queries (SSE etc)`,
				Computed: false,
				Optional: true,
				Required: false,
			},
			"query_type": schema.StringAttribute{
				MarkdownDescription: `Specify the query flavor
TODO make this required and give it a default`,
				Computed: false,
				Optional: true,
				Required: false,
			},

			"rendered_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *QueryGrafanaPyroscopeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryGrafanaPyroscopeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryGrafanaPyroscopeDataSourceModel

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
	data.RenderedJSON = types.StringValue(string(JSONConfig))

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
