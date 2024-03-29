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
	_ datasource.DataSource              = &QueryTestDataDataSource{}
	_ datasource.DataSourceWithConfigure = &QueryTestDataDataSource{}
)

func NewQueryTestDataDataSource() datasource.DataSource {
	return &QueryTestDataDataSource{}
}

// QueryTestDataDataSource defines the data source implementation.
type QueryTestDataDataSource struct{}

type QueryTestDataDataSourceModel_Stream struct {
	Type   types.String `tfsdk:"type"`
	Speed  types.Int64  `tfsdk:"speed"`
	Spread types.Int64  `tfsdk:"spread"`
	Noise  types.Int64  `tfsdk:"noise"`
	Bands  types.Int64  `tfsdk:"bands"`
	Url    types.String `tfsdk:"url"`
}

func (m QueryTestDataDataSourceModel_Stream) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Stream struct {
		Type   string  `json:"type"`
		Speed  int64   `json:"speed"`
		Spread int64   `json:"spread"`
		Noise  int64   `json:"noise"`
		Bands  *int64  `json:"bands,omitempty"`
		Url    *string `json:"url,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_type := m.Type.ValueString()
	attr_speed := m.Speed.ValueInt64()
	attr_spread := m.Spread.ValueInt64()
	attr_noise := m.Noise.ValueInt64()
	attr_bands := m.Bands.ValueInt64Pointer()
	attr_url := m.Url.ValueStringPointer()

	model := &jsonQueryTestDataDataSourceModel_Stream{
		Type:   attr_type,
		Speed:  attr_speed,
		Spread: attr_spread,
		Noise:  attr_noise,
		Bands:  attr_bands,
		Url:    attr_url,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Stream) ApplyDefaults() QueryTestDataDataSourceModel_Stream {

	return m
}

type QueryTestDataDataSourceModel_PulseWave struct {
	TimeStep types.Int64   `tfsdk:"time_step"`
	OnCount  types.Int64   `tfsdk:"on_count"`
	OffCount types.Int64   `tfsdk:"off_count"`
	OnValue  types.Float64 `tfsdk:"on_value"`
	OffValue types.Float64 `tfsdk:"off_value"`
}

func (m QueryTestDataDataSourceModel_PulseWave) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_PulseWave struct {
		TimeStep *int64   `json:"timeStep,omitempty"`
		OnCount  *int64   `json:"onCount,omitempty"`
		OffCount *int64   `json:"offCount,omitempty"`
		OnValue  *float64 `json:"onValue,omitempty"`
		OffValue *float64 `json:"offValue,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_timestep := m.TimeStep.ValueInt64Pointer()
	attr_oncount := m.OnCount.ValueInt64Pointer()
	attr_offcount := m.OffCount.ValueInt64Pointer()
	attr_onvalue := m.OnValue.ValueFloat64Pointer()
	attr_offvalue := m.OffValue.ValueFloat64Pointer()

	model := &jsonQueryTestDataDataSourceModel_PulseWave{
		TimeStep: attr_timestep,
		OnCount:  attr_oncount,
		OffCount: attr_offcount,
		OnValue:  attr_onvalue,
		OffValue: attr_offvalue,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_PulseWave) ApplyDefaults() QueryTestDataDataSourceModel_PulseWave {

	return m
}

type QueryTestDataDataSourceModel_Sim_Key struct {
	Type types.String  `tfsdk:"type"`
	Tick types.Float64 `tfsdk:"tick"`
	Uid  types.String  `tfsdk:"uid"`
}

func (m QueryTestDataDataSourceModel_Sim_Key) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Sim_Key struct {
		Type string  `json:"type"`
		Tick float64 `json:"tick"`
		Uid  *string `json:"uid,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_type := m.Type.ValueString()
	attr_tick := m.Tick.ValueFloat64()
	attr_uid := m.Uid.ValueStringPointer()

	model := &jsonQueryTestDataDataSourceModel_Sim_Key{
		Type: attr_type,
		Tick: attr_tick,
		Uid:  attr_uid,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Sim_Key) ApplyDefaults() QueryTestDataDataSourceModel_Sim_Key {

	return m
}

type QueryTestDataDataSourceModel_Sim_Config struct {
}

func (m QueryTestDataDataSourceModel_Sim_Config) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Sim_Config struct {
	}

	m = m.ApplyDefaults()

	model := &jsonQueryTestDataDataSourceModel_Sim_Config{}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Sim_Config) ApplyDefaults() QueryTestDataDataSourceModel_Sim_Config {

	return m
}

type QueryTestDataDataSourceModel_Sim struct {
	Key    *QueryTestDataDataSourceModel_Sim_Key    `tfsdk:"key"`
	Config *QueryTestDataDataSourceModel_Sim_Config `tfsdk:"config"`
	Stream types.Bool                               `tfsdk:"stream"`
	Last   types.Bool                               `tfsdk:"last"`
}

func (m QueryTestDataDataSourceModel_Sim) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Sim struct {
		Key    interface{} `json:"key,omitempty"`
		Config interface{} `json:"config,omitempty"`
		Stream *bool       `json:"stream,omitempty"`
		Last   *bool       `json:"last,omitempty"`
	}

	m = m.ApplyDefaults()
	var attr_key interface{}
	if m.Key != nil {
		attr_key = m.Key
	}
	var attr_config interface{}
	if m.Config != nil {
		attr_config = m.Config
	}
	attr_stream := m.Stream.ValueBoolPointer()
	attr_last := m.Last.ValueBoolPointer()

	model := &jsonQueryTestDataDataSourceModel_Sim{
		Key:    attr_key,
		Config: attr_config,
		Stream: attr_stream,
		Last:   attr_last,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Sim) ApplyDefaults() QueryTestDataDataSourceModel_Sim {

	return m
}

type QueryTestDataDataSourceModel_CsvWave struct {
	TimeStep  types.Int64  `tfsdk:"time_step"`
	Name      types.String `tfsdk:"name"`
	ValuesCSV types.String `tfsdk:"values_csv"`
	Labels    types.String `tfsdk:"labels"`
}

func (m QueryTestDataDataSourceModel_CsvWave) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_CsvWave struct {
		TimeStep  *int64  `json:"timeStep,omitempty"`
		Name      *string `json:"name,omitempty"`
		ValuesCSV *string `json:"valuesCSV,omitempty"`
		Labels    *string `json:"labels,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_timestep := m.TimeStep.ValueInt64Pointer()
	attr_name := m.Name.ValueStringPointer()
	attr_valuescsv := m.ValuesCSV.ValueStringPointer()
	attr_labels := m.Labels.ValueStringPointer()

	model := &jsonQueryTestDataDataSourceModel_CsvWave{
		TimeStep:  attr_timestep,
		Name:      attr_name,
		ValuesCSV: attr_valuescsv,
		Labels:    attr_labels,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_CsvWave) ApplyDefaults() QueryTestDataDataSourceModel_CsvWave {

	return m
}

type QueryTestDataDataSourceModel_Nodes struct {
	Type  types.String `tfsdk:"type"`
	Count types.Int64  `tfsdk:"count"`
}

func (m QueryTestDataDataSourceModel_Nodes) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Nodes struct {
		Type  *string `json:"type,omitempty"`
		Count *int64  `json:"count,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_type := m.Type.ValueStringPointer()
	attr_count := m.Count.ValueInt64Pointer()

	model := &jsonQueryTestDataDataSourceModel_Nodes{
		Type:  attr_type,
		Count: attr_count,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Nodes) ApplyDefaults() QueryTestDataDataSourceModel_Nodes {

	return m
}

type QueryTestDataDataSourceModel_Usa struct {
	Mode   types.String `tfsdk:"mode"`
	Period types.String `tfsdk:"period"`
	Fields types.List   `tfsdk:"fields"`
	States types.List   `tfsdk:"states"`
}

func (m QueryTestDataDataSourceModel_Usa) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel_Usa struct {
		Mode   *string  `json:"mode,omitempty"`
		Period *string  `json:"period,omitempty"`
		Fields []string `json:"fields,omitempty"`
		States []string `json:"states,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_mode := m.Mode.ValueStringPointer()
	attr_period := m.Period.ValueStringPointer()
	attr_fields := []string{}
	for _, v := range m.Fields.Elements() {
		attr_fields = append(attr_fields, v.(types.String).ValueString())
	}
	attr_states := []string{}
	for _, v := range m.States.Elements() {
		attr_states = append(attr_states, v.(types.String).ValueString())
	}

	model := &jsonQueryTestDataDataSourceModel_Usa{
		Mode:   attr_mode,
		Period: attr_period,
		Fields: attr_fields,
		States: attr_states,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel_Usa) ApplyDefaults() QueryTestDataDataSourceModel_Usa {
	if len(m.Fields.Elements()) == 0 {
		m.Fields, _ = types.ListValue(types.StringType, []attr.Value{})
	}
	if len(m.States.Elements()) == 0 {
		m.States, _ = types.ListValue(types.StringType, []attr.Value{})
	}
	return m
}

type QueryTestDataDataSourceModel struct {
	RenderedJSON    types.String                            `tfsdk:"rendered_json"`
	Alias           types.String                            `tfsdk:"alias"`
	ScenarioId      types.String                            `tfsdk:"scenario_id"`
	StringInput     types.String                            `tfsdk:"string_input"`
	Stream          *QueryTestDataDataSourceModel_Stream    `tfsdk:"stream"`
	PulseWave       *QueryTestDataDataSourceModel_PulseWave `tfsdk:"pulse_wave"`
	Sim             *QueryTestDataDataSourceModel_Sim       `tfsdk:"sim"`
	CsvWave         []QueryTestDataDataSourceModel_CsvWave  `tfsdk:"csv_wave"`
	Labels          types.String                            `tfsdk:"labels"`
	Lines           types.Int64                             `tfsdk:"lines"`
	LevelColumn     types.Bool                              `tfsdk:"level_column"`
	Channel         types.String                            `tfsdk:"channel"`
	Nodes           *QueryTestDataDataSourceModel_Nodes     `tfsdk:"nodes"`
	CsvFileName     types.String                            `tfsdk:"csv_file_name"`
	CsvContent      types.String                            `tfsdk:"csv_content"`
	RawFrameContent types.String                            `tfsdk:"raw_frame_content"`
	SeriesCount     types.Int64                             `tfsdk:"series_count"`
	Usa             *QueryTestDataDataSourceModel_Usa       `tfsdk:"usa"`
	ErrorType       types.String                            `tfsdk:"error_type"`
	SpanCount       types.Int64                             `tfsdk:"span_count"`
	DropPercent     types.Float64                           `tfsdk:"drop_percent"`
	RefId           types.String                            `tfsdk:"ref_id"`
	Hide            types.Bool                              `tfsdk:"hide"`
	QueryType       types.String                            `tfsdk:"query_type"`
}

func (m QueryTestDataDataSourceModel) MarshalJSON() ([]byte, error) {
	type jsonQueryTestDataDataSourceModel struct {
		Alias           *string       `json:"alias,omitempty"`
		ScenarioId      *string       `json:"scenarioId,omitempty"`
		StringInput     *string       `json:"stringInput,omitempty"`
		Stream          interface{}   `json:"stream,omitempty"`
		PulseWave       interface{}   `json:"pulseWave,omitempty"`
		Sim             interface{}   `json:"sim,omitempty"`
		CsvWave         []interface{} `json:"csvWave,omitempty"`
		Labels          *string       `json:"labels,omitempty"`
		Lines           *int64        `json:"lines,omitempty"`
		LevelColumn     *bool         `json:"levelColumn,omitempty"`
		Channel         *string       `json:"channel,omitempty"`
		Nodes           interface{}   `json:"nodes,omitempty"`
		CsvFileName     *string       `json:"csvFileName,omitempty"`
		CsvContent      *string       `json:"csvContent,omitempty"`
		RawFrameContent *string       `json:"rawFrameContent,omitempty"`
		SeriesCount     *int64        `json:"seriesCount,omitempty"`
		Usa             interface{}   `json:"usa,omitempty"`
		ErrorType       *string       `json:"errorType,omitempty"`
		SpanCount       *int64        `json:"spanCount,omitempty"`
		DropPercent     *float64      `json:"dropPercent,omitempty"`
		RefId           string        `json:"refId"`
		Hide            *bool         `json:"hide,omitempty"`
		QueryType       *string       `json:"queryType,omitempty"`
	}

	m = m.ApplyDefaults()
	attr_alias := m.Alias.ValueStringPointer()
	attr_scenarioid := m.ScenarioId.ValueStringPointer()
	attr_stringinput := m.StringInput.ValueStringPointer()
	var attr_stream interface{}
	if m.Stream != nil {
		attr_stream = m.Stream
	}
	var attr_pulsewave interface{}
	if m.PulseWave != nil {
		attr_pulsewave = m.PulseWave
	}
	var attr_sim interface{}
	if m.Sim != nil {
		attr_sim = m.Sim
	}
	attr_csvwave := []interface{}{}
	for _, v := range m.CsvWave {
		attr_csvwave = append(attr_csvwave, v)
	}
	attr_labels := m.Labels.ValueStringPointer()
	attr_lines := m.Lines.ValueInt64Pointer()
	attr_levelcolumn := m.LevelColumn.ValueBoolPointer()
	attr_channel := m.Channel.ValueStringPointer()
	var attr_nodes interface{}
	if m.Nodes != nil {
		attr_nodes = m.Nodes
	}
	attr_csvfilename := m.CsvFileName.ValueStringPointer()
	attr_csvcontent := m.CsvContent.ValueStringPointer()
	attr_rawframecontent := m.RawFrameContent.ValueStringPointer()
	attr_seriescount := m.SeriesCount.ValueInt64Pointer()
	var attr_usa interface{}
	if m.Usa != nil {
		attr_usa = m.Usa
	}
	attr_errortype := m.ErrorType.ValueStringPointer()
	attr_spancount := m.SpanCount.ValueInt64Pointer()
	attr_droppercent := m.DropPercent.ValueFloat64Pointer()
	attr_refid := m.RefId.ValueString()
	attr_hide := m.Hide.ValueBoolPointer()
	attr_querytype := m.QueryType.ValueStringPointer()

	model := &jsonQueryTestDataDataSourceModel{
		Alias:           attr_alias,
		ScenarioId:      attr_scenarioid,
		StringInput:     attr_stringinput,
		Stream:          attr_stream,
		PulseWave:       attr_pulsewave,
		Sim:             attr_sim,
		CsvWave:         attr_csvwave,
		Labels:          attr_labels,
		Lines:           attr_lines,
		LevelColumn:     attr_levelcolumn,
		Channel:         attr_channel,
		Nodes:           attr_nodes,
		CsvFileName:     attr_csvfilename,
		CsvContent:      attr_csvcontent,
		RawFrameContent: attr_rawframecontent,
		SeriesCount:     attr_seriescount,
		Usa:             attr_usa,
		ErrorType:       attr_errortype,
		SpanCount:       attr_spancount,
		DropPercent:     attr_droppercent,
		RefId:           attr_refid,
		Hide:            attr_hide,
		QueryType:       attr_querytype,
	}
	return json.Marshal(model)
}

func (m QueryTestDataDataSourceModel) ApplyDefaults() QueryTestDataDataSourceModel {
	if m.ScenarioId.IsNull() {
		m.ScenarioId = types.StringValue(`random_walk`)
	}
	return m
}

func (d *QueryTestDataDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_query_test_data"
}

func (d *QueryTestDataDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"scenario_id": schema.StringAttribute{
				MarkdownDescription: ` Defaults to "random_walk".`,
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"string_input": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"stream": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"speed": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"spread": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"noise": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            false,
						Required:            true,
					},
					"bands": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"pulse_wave": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"time_step": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"on_count": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"off_count": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"on_value": schema.Float64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"off_value": schema.Float64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"sim": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"key": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"tick": schema.Float64Attribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            false,
								Required:            true,
							},
							"uid": schema.StringAttribute{
								MarkdownDescription: ``,
								Computed:            false,
								Optional:            true,
								Required:            false,
							},
						},
					},
					"config": schema.SingleNestedAttribute{
						MarkdownDescription: ``,
						Computed:            true,
						Optional:            true,
						Required:            false,
					},
					"stream": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"last": schema.BoolAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"csv_wave": schema.ListNestedAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"time_step": schema.Int64Attribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            true,
							Required:            false,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            true,
							Required:            false,
						},
						"values_csv": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            true,
							Required:            false,
						},
						"labels": schema.StringAttribute{
							MarkdownDescription: ``,
							Computed:            false,
							Optional:            true,
							Required:            false,
						},
					},
				},
			},
			"labels": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"lines": schema.Int64Attribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"level_column": schema.BoolAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"channel": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"nodes": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"count": schema.Int64Attribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
				},
			},
			"csv_file_name": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"csv_content": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"raw_frame_content": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"series_count": schema.Int64Attribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"usa": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Optional:            true,
				Required:            false,
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"period": schema.StringAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
					},
					"fields": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						ElementType:         types.StringType,
					},
					"states": schema.ListAttribute{
						MarkdownDescription: ``,
						Computed:            false,
						Optional:            true,
						Required:            false,
						ElementType:         types.StringType,
					},
				},
			},
			"error_type": schema.StringAttribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"span_count": schema.Int64Attribute{
				MarkdownDescription: ``,
				Computed:            false,
				Optional:            true,
				Required:            false,
			},
			"drop_percent": schema.Float64Attribute{
				MarkdownDescription: `Drop percentage (the chance we will lose a point 0-100)`,
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

func (d *QueryTestDataDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *QueryTestDataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QueryTestDataDataSourceModel

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
