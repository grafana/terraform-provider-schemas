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
	_ datasource.DataSource              = &CoreRoleDataSource{}
	_ datasource.DataSourceWithConfigure = &CoreRoleDataSource{}
)

func NewCoreRoleDataSource() datasource.DataSource {
	return &CoreRoleDataSource{}
}

// CoreRoleDataSource defines the data source implementation.
type CoreRoleDataSource struct{}

type CoreRoleDataSourceModel struct {
	ToJSON      types.String `tfsdk:"to_json"`
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
	GroupName   types.String `tfsdk:"group_name"`
	Description types.String `tfsdk:"description"`
	Hidden      types.Bool   `tfsdk:"hidden"`
}

func (m CoreRoleDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonCoreRoleDataSourceModel struct {
		Name        string  `json:"name"`
		DisplayName *string `json:"displayName,omitempty"`
		GroupName   *string `json:"groupName,omitempty"`
		Description *string `json:"description,omitempty"`
		Hidden      bool    `json:"hidden"`
	}

	m = m.ApplyDefaults()
	attr_name := m.Name.ValueString()
	attr_displayname := m.DisplayName.ValueString()
	attr_groupname := m.GroupName.ValueString()
	attr_description := m.Description.ValueString()
	attr_hidden := m.Hidden.ValueBool()

	model := &jsonCoreRoleDataSourceModel{
		Name:        attr_name,
		DisplayName: &attr_displayname,
		GroupName:   &attr_groupname,
		Description: &attr_description,
		Hidden:      attr_hidden,
	}
	return json.Marshal(model)
}

func (m CoreRoleDataSourceModel) ApplyDefaults() CoreRoleDataSourceModel {

	return m
}

func (d *CoreRoleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_role"
}

func (d *CoreRoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: `The role identifier managed:builtins:editor:permissions`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"display_name": schema.StringAttribute{
				MarkdownDescription: `Optional display`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"group_name": schema.StringAttribute{
				MarkdownDescription: `Name of the team.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Role description`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"hidden": schema.BoolAttribute{
				MarkdownDescription: `Do not show this role`,
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

func (d *CoreRoleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CoreRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreRoleDataSourceModel

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