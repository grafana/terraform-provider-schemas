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
	_ datasource.DataSource              = &QueryElasticsearchDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryElasticsearchDataSource{}
)

func NewQueryElasticsearchDataSource() datasource.DataSource {
	return &QueryElasticsearchDataSource{}
}

// QueryElasticsearchDataSource defines the data source implementation.
type QueryElasticsearchDataSource struct{}

type QueryElasticsearchDataSourceModel struct {
	RenderedJSON types.String `tfsdk:"rendered_json"`
	Alias        types.String `tfsdk:"alias"`
	Query        types.String `tfsdk:"query"`
	TimeField    types.String `tfsdk:"time_field"`
	RefId        types.String `tfsdk:"ref_id"`
	Hide         types.Bool   `tfsdk:"hide"`
	QueryType    types.String `tfsdk:"query_type"`
}

func (m QueryElasticsearchDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryElasticsearchDataSourceModel struct {
		Alias     *string `json:"alias,omitempty"`
		Query     *string `json:"query,omitempty"`
		TimeField *string `json:"timeField,omitempty"`
		RefId     string  `json:"refId"`
		Hide      *bool   `json:"hide,omitempty"`
		QueryType *string `json:"queryType,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_alias := m.Alias.ValueString()
	attr_query := m.Query.ValueString()
	attr_timefield := m.TimeField.ValueString()
	attr_refid := m.RefId.ValueString()
	attr_hide := m.Hide.ValueBool()
	attr_querytype := m.QueryType.ValueString()

	model := &jsonQueryElasticsearchDataSourceModel{
		Alias:     &attr_alias,
		Query:     &attr_query,
		TimeField: &attr_timefield,
		RefId:     attr_refid,
		Hide:      &attr_hide,
		QueryType: &attr_querytype,
	}
	return json.Marshal(model)
}

func (m QueryElasticsearchDataSourceModel) ApplyDefaults() QueryElasticsearchDataSourceModel {

	return m
}

func (d *QueryElasticsearchDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_elasticsearch"
}

func (d *QueryElasticsearchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				MarkdownDescription: `Alias pattern`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: `Lucene query`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"time_field": schema.StringAttribute{
				MarkdownDescription: `Name of time field`,
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

func (d *QueryElasticsearchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryElasticsearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryElasticsearchDataSourceModel

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
