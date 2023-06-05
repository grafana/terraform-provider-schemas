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
	_ datasource.DataSource              = &QueryGoogleCloudMonitoringDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryGoogleCloudMonitoringDataSource{}
)

func NewQueryGoogleCloudMonitoringDataSource() datasource.DataSource {
	return &QueryGoogleCloudMonitoringDataSource{}
}

// QueryGoogleCloudMonitoringDataSource defines the data source implementation.
type QueryGoogleCloudMonitoringDataSource struct{}

type QueryGoogleCloudMonitoringDataSourceModel struct {
	ToJSON types.String `tfsdk:"to_json"`
}

func (m QueryGoogleCloudMonitoringDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryGoogleCloudMonitoringDataSourceModel struct {
	}

	m = m.ApplyDefaults()

	model := &jsonQueryGoogleCloudMonitoringDataSourceModel{}
	return json.Marshal(model)
}

func (m QueryGoogleCloudMonitoringDataSourceModel) ApplyDefaults() QueryGoogleCloudMonitoringDataSourceModel {

	return m
}

func (d *QueryGoogleCloudMonitoringDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_google_cloud_monitoring"
}

func (d *QueryGoogleCloudMonitoringDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *QueryGoogleCloudMonitoringDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryGoogleCloudMonitoringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryGoogleCloudMonitoringDataSourceModel

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
