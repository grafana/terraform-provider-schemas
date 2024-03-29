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
	_ datasource.DataSource              = &QueryLokiDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryLokiDataSource{}
)

func NewQueryLokiDataSource() datasource.DataSource {
	return &QueryLokiDataSource{}
}

// QueryLokiDataSource defines the data source implementation.
type QueryLokiDataSource struct{}

type QueryLokiDataSourceModel struct {
	RenderedJSON types.String `tfsdk:"rendered_json"`
	Expr         types.String `tfsdk:"expr"`
	LegendFormat types.String `tfsdk:"legend_format"`
	MaxLines     types.Int64  `tfsdk:"max_lines"`
	Resolution   types.Int64  `tfsdk:"resolution"`
	EditorMode   types.String `tfsdk:"editor_mode"`
	Range        types.Bool   `tfsdk:"range"`
	Instant      types.Bool   `tfsdk:"instant"`
	Step         types.String `tfsdk:"step"`
	RefId        types.String `tfsdk:"ref_id"`
	Hide         types.Bool   `tfsdk:"hide"`
	QueryType    types.String `tfsdk:"query_type"`
}

func (m QueryLokiDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryLokiDataSourceModel struct {
		Expr         string  `json:"expr"`
		LegendFormat *string `json:"legendFormat,omitempty"`
		MaxLines     *int64  `json:"maxLines,omitempty"`
		Resolution   *int64  `json:"resolution,omitempty"`
		EditorMode   *string `json:"editorMode,omitempty"`
		Range        *bool   `json:"range,omitempty"`
		Instant      *bool   `json:"instant,omitempty"`
		Step         *string `json:"step,omitempty"`
		RefId        string  `json:"refId"`
		Hide         *bool   `json:"hide,omitempty"`
		QueryType    *string `json:"queryType,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_expr := m.Expr.ValueString()
	attr_legendformat := m.LegendFormat.ValueStringPointer()
	attr_maxlines := m.MaxLines.ValueInt64Pointer()
	attr_resolution := m.Resolution.ValueInt64Pointer()
	attr_editormode := m.EditorMode.ValueStringPointer()
	attr_range := m.Range.ValueBoolPointer()
	attr_instant := m.Instant.ValueBoolPointer()
	attr_step := m.Step.ValueStringPointer()
	attr_refid := m.RefId.ValueString()
	attr_hide := m.Hide.ValueBoolPointer()
	attr_querytype := m.QueryType.ValueStringPointer()

	model := &jsonQueryLokiDataSourceModel{
		Expr:         attr_expr,
		LegendFormat: attr_legendformat,
		MaxLines:     attr_maxlines,
		Resolution:   attr_resolution,
		EditorMode:   attr_editormode,
		Range:        attr_range,
		Instant:      attr_instant,
		Step:         attr_step,
		RefId:        attr_refid,
		Hide:         attr_hide,
		QueryType:    attr_querytype,
	}
	return json.Marshal(model)
}

func (m QueryLokiDataSourceModel) ApplyDefaults() QueryLokiDataSourceModel {

	return m
}

func (d *QueryLokiDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_loki"
}

func (d *QueryLokiDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"expr": schema.StringAttribute{
				MarkdownDescription: `The LogQL query.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"legend_format": schema.StringAttribute{
				MarkdownDescription: `Used to override the name of the series.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"max_lines": schema.Int64Attribute{
				MarkdownDescription: `Used to limit the number of log rows returned.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `@deprecated, now use step.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				DeprecationMessage:  `Now use step.`,
			},
			"editor_mode": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"range": schema.BoolAttribute{
				MarkdownDescription: `@deprecated, now use queryType.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				DeprecationMessage:  `Now use queryType.`,
			},
			"instant": schema.BoolAttribute{
				MarkdownDescription: `@deprecated, now use queryType.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
				DeprecationMessage:  `Now use queryType.`,
			},
			"step": schema.StringAttribute{
				MarkdownDescription: `Used to set step value for range queries.`,
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

func (d *QueryLokiDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryLokiDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryLokiDataSourceModel

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
