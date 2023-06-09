package terraform

import (
	"fmt"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/codegen"
	"github.com/grafana/terraform-provider-schemas/gen/terraform/cuetf"
	"golang.org/x/tools/imports"
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
	fname := "datasource_" + name + "_gen.go"

	bytes, err = imports.Process(fname, bytes, nil)
	if err != nil {
		return nil, fmt.Errorf("goimports processing of generated file failed: %w", err)
	}

	return codejen.NewFile(fname, bytes, j), nil
}
