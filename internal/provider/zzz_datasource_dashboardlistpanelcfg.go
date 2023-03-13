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
	_ datasource.DataSource              = &DashboardListPanelCfgDataSource{}
	_ datasource.DataSourceWithConfigure = &DashboardListPanelCfgDataSource{}
)

func NewDashboardListPanelCfgDataSource() datasource.DataSource {
	return &DashboardListPanelCfgDataSource{}
}

// DashboardListPanelCfgDataSource defines the data source implementation.
type DashboardListPanelCfgDataSource struct{}

// DashboardListPanelCfgDataSourceModel describes the data source data model.
type DashboardListPanelCfgDataSourceModel struct {
	PanelLayout  types.String `tfsdk:"panel_layout" json:"PanelLayout"`
	PanelOptions *struct {
		Layout             types.String `tfsdk:"layout" json:"layout"`
		ShowStarred        types.Bool   `tfsdk:"show_starred" json:"showStarred"`
		ShowRecentlyViewed types.Bool   `tfsdk:"show_recently_viewed" json:"showRecentlyViewed"`
		ShowSearch         types.Bool   `tfsdk:"show_search" json:"showSearch"`
		ShowHeadings       types.Bool   `tfsdk:"show_headings" json:"showHeadings"`
		MaxItems           types.Int64  `tfsdk:"max_items" json:"maxItems"`
		Query              types.String `tfsdk:"query" json:"query"`
		FolderId           types.Int64  `tfsdk:"folder_id" json:"folderId"`
		Tags               types.List   `tfsdk:"tags" json:"tags"`
	} `tfsdk:"panel_options" json:"PanelOptions"`
	ToJSON types.String `tfsdk:"to_json"`
}

func (d *DashboardListPanelCfgDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dashboardlistpanelcfg"
}

func (d *DashboardListPanelCfgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"panel_layout": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},

			"panel_options": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            false,
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"layout": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},

					"show_starred": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"show_recently_viewed": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"show_search": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"show_headings": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"max_items": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"query": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},

					"folder_id": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
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
				},
			},

			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *DashboardListPanelCfgDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *DashboardListPanelCfgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DashboardListPanelCfgDataSourceModel

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