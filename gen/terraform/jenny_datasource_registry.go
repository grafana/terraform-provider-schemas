package terraform

import (
	"fmt"
	"strings"

	"github.com/grafana/codejen"
	"github.com/grafana/kindsys"
	"github.com/grafana/terraform-provider-schemas/gen/terraform/cuetf"
)

var datasources []string

type TerraformCoreRegistryJenny struct {
}

func (j TerraformCoreRegistryJenny) JennyName() string {
	return "TerraformCoreRegistryJenny"
}

func (j TerraformCoreRegistryJenny) Generate(k ...kindsys.Kind) (*codejen.File, error) {
	for _, k := range k {
		datasources = append(datasources, cuetf.GetStructName(k.Lineage().Name()))
	}

	return nil, nil
}

type TerraformComposableRegistryJenny struct {
}

func (j TerraformComposableRegistryJenny) JennyName() string {
	return "TerraformComposableRegistryJenny"
}

func (j TerraformComposableRegistryJenny) Generate(k ...kindsys.Composable) (*codejen.File, error) {
	for _, k := range k {
		datasources = append(datasources, cuetf.GetStructName(k.Lineage().Name()))
	}

	datasourceConstructors := []string{}
	for _, datasource := range datasources {
		datasourceConstructors = append(datasourceConstructors, "New"+datasource)
	}

	bytes := []byte(fmt.Sprintf(`package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var datasources = []func() datasource.DataSource{
	%s,
}
`, strings.Join(datasourceConstructors, ",\n	")))

	return codejen.NewFile("registry_gen.go", bytes, j), nil
}
