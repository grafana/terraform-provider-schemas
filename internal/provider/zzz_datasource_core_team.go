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
	_ datasource.DataSource              = &CoreTeamDataSource{}
	_ datasource.DataSourceWithConfigure = &CoreTeamDataSource{}
)

func NewCoreTeamDataSource() datasource.DataSource {
	return &CoreTeamDataSource{}
}

// CoreTeamDataSource defines the data source implementation.
type CoreTeamDataSource struct{}

// CoreTeamDataSourceModel describes the data source data model.
type CoreTeamDataSourceModel struct {
	OrgId         types.Int64  `tfsdk:"org_id"`
	Name          types.String `tfsdk:"name"`
	Email         types.String `tfsdk:"email"`
	AvatarUrl     types.String `tfsdk:"avatar_url"`
	MemberCount   types.Int64  `tfsdk:"member_count"`
	Permission    types.Int64  `tfsdk:"permission"`
	AccessControl *struct {
	} `tfsdk:"access_control"`
	Created types.Int64  `tfsdk:"created"`
	Updated types.Int64  `tfsdk:"updated"`
	ToJSON  types.String `tfsdk:"to_json"`
}

// CoreTeamDataSourceModelJSON describes the data source data model when exported to json.
type CoreTeamDataSourceModelJSON struct {
	OrgId         int64   `json:"orgId"`
	Name          string  `json:"name"`
	Email         *string `json:"email,omitempty"`
	AvatarUrl     *string `json:"avatarUrl,omitempty"`
	MemberCount   int64   `json:"memberCount"`
	Permission    int64   `json:"permission"`
	AccessControl *struct {
	} `json:"accessControl,omitempty"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}

func (d *CoreTeamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_team"
}

func (d *CoreTeamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.Int64Attribute{
				MarkdownDescription: `OrgId is the ID of an organisation the team belongs to.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the team.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: `Email of the team.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"avatar_url": schema.StringAttribute{
				MarkdownDescription: `AvatarUrl is the team's avatar URL.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"member_count": schema.Int64Attribute{
				MarkdownDescription: `MemberCount is the number of the team members.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"permission": schema.Int64Attribute{
				MarkdownDescription: `TODO - it seems it's a team_member.permission, unlikely it should belong to the team kind`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"access_control": schema.SingleNestedAttribute{
				MarkdownDescription: `AccessControl metadata associated with a given resource.`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"created": schema.Int64Attribute{
				MarkdownDescription: `Created indicates when the team was created.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"updated": schema.Int64Attribute{
				MarkdownDescription: `Updated indicates when the team was updated.`,
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

func (d *CoreTeamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CoreTeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreTeamDataSourceModel

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

func (d CoreTeamDataSourceModel) MarshalJSON() ([]byte, error) {
	attr_orgid := d.OrgId.ValueInt64()
	attr_name := d.Name.ValueString()
	attr_email := d.Email.ValueString()
	attr_avatarurl := d.AvatarUrl.ValueString()
	attr_membercount := d.MemberCount.ValueInt64()
	attr_permission := d.Permission.ValueInt64()
	attr_created := d.Created.ValueInt64()
	attr_updated := d.Updated.ValueInt64()

	model := &CoreTeamDataSourceModelJSON{
		OrgId:       attr_orgid,
		Name:        attr_name,
		Email:       &attr_email,
		AvatarUrl:   &attr_avatarurl,
		MemberCount: attr_membercount,
		Permission:  attr_permission,
		Created:     attr_created,
		Updated:     attr_updated,
	}
	return json.Marshal(model)
}
