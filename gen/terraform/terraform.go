package terraform

import (
	jen "github.com/grafana/terraform-provider-schemas/gen/internal"
)

func JenniesForTerraform() jen.TargetJennies {
	tgt := jen.NewTargetJennies()

	tgt.Core.Append(
		jen.LatestJenny("", TerraformDataSourceJenny{}),
		&TerraformCoreRegistryJenny{},
	)

	tgt.Composable.Append(
		jen.ComposableLatestMajorsOrXJenny("", true, TerraformDataSourceJenny{}),
		&TerraformComposableRegistryJenny{},
	)

	return tgt
}
