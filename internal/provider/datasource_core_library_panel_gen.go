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
	_ datasource.DataSource              = &CoreLibraryPanelDataSource{}
	_ datasource.DataSourceWithConfigure = &CoreLibraryPanelDataSource{}
)

func NewCoreLibraryPanelDataSource() datasource.DataSource {
	return &CoreLibraryPanelDataSource{}
}

// CoreLibraryPanelDataSource defines the data source implementation.
type CoreLibraryPanelDataSource struct{}

type CoreLibraryPanelDataSourceModel_Model struct {
}

func (m CoreLibraryPanelDataSourceModel_Model) MarshalJSON() ([]byte, error) {
	type jsonCoreLibraryPanelDataSourceModel_Model struct {
	}

	m = m.ApplyDefaults()

	model := &jsonCoreLibraryPanelDataSourceModel_Model{}
	return json.Marshal(model)
}

func (m CoreLibraryPanelDataSourceModel_Model) ApplyDefaults() CoreLibraryPanelDataSourceModel_Model {

	return m
}

type CoreLibraryPanelDataSourceModel_Meta_CreatedBy struct {
	Name      types.String `tfsdk:"name"`
	AvatarUrl types.String `tfsdk:"avatar_url"`
}

func (m CoreLibraryPanelDataSourceModel_Meta_CreatedBy) MarshalJSON() ([]byte, error) {
	type jsonCoreLibraryPanelDataSourceModel_Meta_CreatedBy struct {
		Name      string `json:"name"`
		AvatarUrl string `json:"avatarUrl"`
	}

	m = m.ApplyDefaults()
	attr_name := m.Name.ValueString()
	attr_avatarurl := m.AvatarUrl.ValueString()

	model := &jsonCoreLibraryPanelDataSourceModel_Meta_CreatedBy{
		Name:      attr_name,
		AvatarUrl: attr_avatarurl,
	}
	return json.Marshal(model)
}

func (m CoreLibraryPanelDataSourceModel_Meta_CreatedBy) ApplyDefaults() CoreLibraryPanelDataSourceModel_Meta_CreatedBy {

	return m
}

type CoreLibraryPanelDataSourceModel_Meta_UpdatedBy struct {
	Name      types.String `tfsdk:"name"`
	AvatarUrl types.String `tfsdk:"avatar_url"`
}

func (m CoreLibraryPanelDataSourceModel_Meta_UpdatedBy) MarshalJSON() ([]byte, error) {
	type jsonCoreLibraryPanelDataSourceModel_Meta_UpdatedBy struct {
		Name      string `json:"name"`
		AvatarUrl string `json:"avatarUrl"`
	}

	m = m.ApplyDefaults()
	attr_name := m.Name.ValueString()
	attr_avatarurl := m.AvatarUrl.ValueString()

	model := &jsonCoreLibraryPanelDataSourceModel_Meta_UpdatedBy{
		Name:      attr_name,
		AvatarUrl: attr_avatarurl,
	}
	return json.Marshal(model)
}

func (m CoreLibraryPanelDataSourceModel_Meta_UpdatedBy) ApplyDefaults() CoreLibraryPanelDataSourceModel_Meta_UpdatedBy {

	return m
}

type CoreLibraryPanelDataSourceModel_Meta struct {
	FolderName          types.String                                    `tfsdk:"folder_name"`
	FolderUid           types.String                                    `tfsdk:"folder_uid"`
	ConnectedDashboards types.Int64                                     `tfsdk:"connected_dashboards"`
	Created             types.String                                    `tfsdk:"created"`
	Updated             types.String                                    `tfsdk:"updated"`
	CreatedBy           *CoreLibraryPanelDataSourceModel_Meta_CreatedBy `tfsdk:"created_by"`
	UpdatedBy           *CoreLibraryPanelDataSourceModel_Meta_UpdatedBy `tfsdk:"updated_by"`
}

func (m CoreLibraryPanelDataSourceModel_Meta) MarshalJSON() ([]byte, error) {
	type jsonCoreLibraryPanelDataSourceModel_Meta struct {
		FolderName          string      `json:"folderName"`
		FolderUid           string      `json:"folderUid"`
		ConnectedDashboards int64       `json:"connectedDashboards"`
		Created             string      `json:"created"`
		Updated             string      `json:"updated"`
		CreatedBy           interface{} `json:"createdBy,omitempty"`
		UpdatedBy           interface{} `json:"updatedBy,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_foldername := m.FolderName.ValueString()
	attr_folderuid := m.FolderUid.ValueString()
	attr_connecteddashboards := m.ConnectedDashboards.ValueInt64()
	attr_created := m.Created.ValueString()
	attr_updated := m.Updated.ValueString()
	var attr_createdby interface{}
	if m.CreatedBy != nil {
		attr_createdby = m.CreatedBy
	}
	var attr_updatedby interface{}
	if m.UpdatedBy != nil {
		attr_updatedby = m.UpdatedBy
	}

	model := &jsonCoreLibraryPanelDataSourceModel_Meta{
		FolderName:          attr_foldername,
		FolderUid:           attr_folderuid,
		ConnectedDashboards: attr_connecteddashboards,
		Created:             attr_created,
		Updated:             attr_updated,
		CreatedBy:           attr_createdby,
		UpdatedBy:           attr_updatedby,
	}
	return json.Marshal(model)
}

func (m CoreLibraryPanelDataSourceModel_Meta) ApplyDefaults() CoreLibraryPanelDataSourceModel_Meta {

	return m
}

type CoreLibraryPanelDataSourceModel struct {
	ToJSON        types.String                           `tfsdk:"to_json"`
	FolderUid     types.String                           `tfsdk:"folder_uid"`
	Uid           types.String                           `tfsdk:"uid"`
	Name          types.String                           `tfsdk:"name"`
	Description   types.String                           `tfsdk:"description"`
	Type          types.String                           `tfsdk:"type"`
	SchemaVersion types.Int64                            `tfsdk:"schema_version"`
	Version       types.Int64                            `tfsdk:"version"`
	Model         *CoreLibraryPanelDataSourceModel_Model `tfsdk:"model"`
	Meta          *CoreLibraryPanelDataSourceModel_Meta  `tfsdk:"meta"`
}

func (m CoreLibraryPanelDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonCoreLibraryPanelDataSourceModel struct {
		FolderUid     *string     `json:"folderUid,omitempty"`
		Uid           string      `json:"uid"`
		Name          string      `json:"name"`
		Description   *string     `json:"description,omitempty"`
		Type          string      `json:"type"`
		SchemaVersion *int64      `json:"schemaVersion,omitempty"`
		Version       int64       `json:"version"`
		Model         interface{} `json:"model,omitempty"`
		Meta          interface{} `json:"meta,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_folderuid := m.FolderUid.ValueString()
	attr_uid := m.Uid.ValueString()
	attr_name := m.Name.ValueString()
	attr_description := m.Description.ValueString()
	attr_type := m.Type.ValueString()
	attr_schemaversion := m.SchemaVersion.ValueInt64()
	attr_version := m.Version.ValueInt64()
	var attr_model interface{}
	if m.Model != nil {
		attr_model = m.Model
	}
	var attr_meta interface{}
	if m.Meta != nil {
		attr_meta = m.Meta
	}

	model := &jsonCoreLibraryPanelDataSourceModel{
		FolderUid:     &attr_folderuid,
		Uid:           attr_uid,
		Name:          attr_name,
		Description:   &attr_description,
		Type:          attr_type,
		SchemaVersion: &attr_schemaversion,
		Version:       attr_version,
		Model:         attr_model,
		Meta:          attr_meta,
	}
	return json.Marshal(model)
}

func (m CoreLibraryPanelDataSourceModel) ApplyDefaults() CoreLibraryPanelDataSourceModel {

	return m
}

func (d *CoreLibraryPanelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_library_panel"
}

func (d *CoreLibraryPanelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"folder_uid": schema.StringAttribute{
				MarkdownDescription: `Folder UID`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"uid": schema.StringAttribute{
				MarkdownDescription: `Library element UID`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Panel name (also saved in the model)`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Panel description`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The panel type (from inside the model)`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"schema_version": schema.Int64Attribute{
				MarkdownDescription: `Dashboard version when this was saved (zero if unknown)`,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: `panel version, incremented each time the dashboard is updated.`,
				Computed:            false,
				Optional:            false,
				Required:            true,
			},
			"model": schema.SingleNestedAttribute{
				MarkdownDescription: `TODO: should be the same panel schema defined in dashboard
Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;`,
				Computed: true,
				Optional: true,
				Required: false,
			},
			"meta": schema.SingleNestedAttribute{
				MarkdownDescription: `Object storage metadata`,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"folder_name": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"folder_uid": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"connected_dashboards": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"created": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"updated": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"created_by": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"avatar_url": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
						},
					},
					"updated_by": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"avatar_url": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
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

func (d *CoreLibraryPanelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *CoreLibraryPanelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreLibraryPanelDataSourceModel

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