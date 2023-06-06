// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformComposableRegistryJenny
//
// Run 'go generate ./' from repository root to regenerate.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var datasources = []func() datasource.DataSource{
	NewCoreAccessPolicyDataSource,
	NewCoreDashboardDataSource,
	NewCoreFolderDataSource,
	NewCoreLibraryPanelDataSource,
	NewCorePlaylistDataSource,
	NewCorePreferencesDataSource,
	NewCorePublicDashboardDataSource,
	NewCoreRoleDataSource,
	NewCoreRoleBindingDataSource,
	NewCoreTeamDataSource,
	NewQueryAzureMonitorDataSource,
	NewQueryGoogleCloudMonitoringDataSource,
	NewQueryCloudWatchDataSource,
	NewQueryElasticsearchDataSource,
	NewQueryLokiDataSource,
	NewQueryParcaDataSource,
	NewQueryGrafanaPyroscopeDataSource,
	NewQueryPrometheusDataSource,
	NewQueryTempoDataSource,
	NewQueryTestDataDataSource,
	NewPanelAlertGroupsDataSource,
	NewPanelAnnotationsListDataSource,
	NewPanelBarChartDataSource,
	NewPanelBarGaugeDataSource,
	NewPanelDashboardListDataSource,
	NewPanelDatagridDataSource,
	NewPanelDebugDataSource,
	NewPanelGaugeDataSource,
	NewPanelGeomapDataSource,
	NewPanelHeatmapDataSource,
	NewPanelHistogramDataSource,
	NewPanelLogsDataSource,
	NewPanelNewsDataSource,
	NewPanelNodeGraphDataSource,
	NewPanelPieChartDataSource,
	NewPanelStatDataSource,
	NewPanelStateTimelineDataSource,
	NewPanelStatusHistoryDataSource,
	NewPanelTableDataSource,
	NewPanelTextDataSource,
	NewPanelTimeSeriesDataSource,
	NewPanelTrendDataSource,
	NewPanelXYChartDataSource,
}