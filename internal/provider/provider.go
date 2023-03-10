package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure GrafanaSchemasProvider satisfies various provider interfaces.
var _ provider.Provider = &GrafanaSchemasProvider{}

// GrafanaSchemasProvider defines the provider implementation.
type GrafanaSchemasProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// GrafanaSchemasProviderModel describes the provider data model.
type GrafanaSchemasProviderModel struct {
}

func (p *GrafanaSchemasProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "grafanaschemas"
	resp.Version = p.version
}

func (p *GrafanaSchemasProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *GrafanaSchemasProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data GrafanaSchemasProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *GrafanaSchemasProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *GrafanaSchemasProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewdashboardDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GrafanaSchemasProvider{
			version: version,
		}
	}
}
