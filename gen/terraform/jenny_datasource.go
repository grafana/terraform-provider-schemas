package terraform

import (
	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/terraform-provider-schemas/gen/terraform/cuetf"
)

type TerraformDataSourceJenny struct{}

func (j TerraformDataSourceJenny) JennyName() string {
	return "TerraformDataSourceJenny"
}

func (j TerraformDataSourceJenny) Generate(sfg codegen.SchemaForGen) (*codejen.File, error) {
	bytes, err := cuetf.GenerateDataSource(sfg.Schema)
	if err != nil {
		return nil, err
	}

	name := cuetf.GetResourceName(sfg.Schema.Lineage().Name())
	return codejen.NewFile("datasource_"+name+"_gen.go", bytes, j), nil
}
