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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &CorePublicDashboardDataSource{}
	_ datasource.DataSourceWithConfigure = &CorePublicDashboardDataSource{}
)

func NewCorePublicDashboardDataSource() datasource.DataSource {
	return &CorePublicDashboardDataSource{}
}

// CorePublicDashboardDataSource defines the data source implementation.
type CorePublicDashboardDataSource struct{}

// CorePublicDashboardDataSourceModel describes the data source data model.
type CorePublicDashboardDataSourceModel struct {
	Uid                  types.String `tfsdk:"uid" json:"uid"`
	DashboardUid         types.String `tfsdk:"dashboard_uid" json:"dashboardUid"`
	AccessToken          types.String `tfsdk:"access_token" json:"accessToken"`
	IsEnabled            types.Bool   `tfsdk:"is_enabled" json:"isEnabled"`
	AnnotationsEnabled   types.Bool   `tfsdk:"annotations_enabled" json:"annotationsEnabled"`
	TimeSelectionEnabled types.Bool   `tfsdk:"time_selection_enabled" json:"timeSelectionEnabled"`
	ToJSON               types.String `tfsdk:"to_json"`
}

func (d *CorePublicDashboardDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_public_dashboard"
}

func (d *CorePublicDashboardDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",
		Attributes: map[string]schema.Attribute{
			"uid": schema.StringAttribute{
				MarkdownDescription: `Unique public dashboard identifier`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"dashboard_uid": schema.StringAttribute{
				MarkdownDescription: `Dashboard unique identifier referenced by this public dashboard`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"access_token": schema.StringAttribute{
				MarkdownDescription: `Unique public access token`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"is_enabled": schema.BoolAttribute{
				MarkdownDescription: `Flag that indicates if the public dashboard is enabled`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"annotations_enabled": schema.BoolAttribute{
				MarkdownDescription: `Flag that indicates if annotations are enabled`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"time_selection_enabled": schema.BoolAttribute{
				MarkdownDescription: `Flag that indicates if the time range picker is enabled`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *CorePublicDashboardDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CorePublicDashboardDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CorePublicDashboardDataSourceModel

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