package types

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/terraform-provider-schemas/gen/terraform/cuetf/internal/utils"
)

type Model struct {
	Name          string
	IsDisjunction bool
	Nodes         []Node
	Nested        bool
}

// terraformModel generates the Terraform SDK model.
func (s *Model) terraformModel() string {
	fields := make([]string, 0)
	if !s.Nested {
		fields = append(fields, "RenderedJSON types.String `tfsdk:\"rendered_json\"`")
	}

	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}
		fields = append(fields, node.TerraformModelField(s.Name))
	}

	return fmt.Sprintf(`type %s struct {
	%s
}
`, s.Name, strings.Join(fields, "\n"))
}

// jsonModel generates the JSON model used to convert the Terraform SDK model to JSON.
func (s *Model) jsonModel() string {
	fields := make([]string, 0)
	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}
		fields = append(fields, node.JSONModelField())
	}

	return fmt.Sprintf(`type json%s struct {
	%s
}
`, s.Name, strings.Join(fields, "\n"))
}

// generateToJSONFunction generates a function that converts the Terraform SDK model to the JSON model representation.
func (s *Model) generateToJSONFunction() string {
	b := strings.Builder{}

	for _, node := range s.Nodes {
		if len(node.DisjunctionKinds) > 0 {
			b.WriteString(s.generateGetAttrFunction(node))
		}
	}

	fmt.Fprintf(&b, "func (m %s) MarshalJSON() ([]byte, error) {\n", s.Name)
	if s.IsDisjunction {
		fmt.Fprintf(&b, "var json_%s interface{}\n", s.Name)
	} else {
		b.WriteString(s.jsonModel() + "\n")
	}
	b.WriteString("m = m.ApplyDefaults()\n")

	structLines := make([]string, 0)
	for _, node := range s.Nodes {
		if !node.IsGenerated() {
			continue
		}

		identifier := "attr_" + strings.ToLower(node.Name)
		if s.IsDisjunction {
			identifier = "json_" + s.Name
		}
		funcString := node.terraformFunc()

		if node.Kind == cue.ListKind {
			subType := node.SubTerraformType()
			subTypeGolang := node.subGolangType()
			subTypeFunc := node.subTerraformFunc()
			if subType != "" {
				fmt.Fprintf(&b, "	%s := []%s{}\n", identifier, subTypeGolang)
				fmt.Fprintf(&b, "	for _, v := range m.%s.Elements() {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = append(%s, v.(types.%s).%s)\n", identifier, identifier, subType, subTypeFunc)
				b.WriteString("	}\n")
			} else if node.SubKind == cue.StructKind {
				fmt.Fprintf(&b, "	%s := []interface{}{}\n", identifier)
				fmt.Fprintf(&b, "	for _, v := range m.%s {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = append(%s, v)\n", identifier, identifier)
				b.WriteString("	}\n")
			}
		} else if node.Kind == cue.StructKind {
			if !s.IsDisjunction {
				fmt.Fprintf(&b, "	var %s interface{}\n", identifier)
			}
			if node.Optional {
				fmt.Fprintf(&b, "	if m.%s != nil {\n", utils.ToCamelCase(node.Name))
				fmt.Fprintf(&b, "		%s = m.%s\n", identifier, utils.ToCamelCase(node.Name))
				b.WriteString("	}\n")
			} else {
				fmt.Fprintf(&b, "	%s = m.%s\n", identifier, utils.ToCamelCase(node.Name))
			}
		} else if len(node.DisjunctionKinds) > 0 {
			fmt.Fprintf(&b, "   %s := m.GetAttr%s()\n", identifier, utils.ToCamelCase(node.Name))
		} else if funcString != "" {
			fmt.Fprintf(&b, "	%s := m.%s.%s\n", identifier, utils.ToCamelCase(node.Name), funcString)
			if node.Optional {
				identifier = "&" + identifier
			}
		}

		structLines = append(structLines, fmt.Sprintf("		%s: %s,\n", utils.ToCamelCase(node.Name), identifier))
	}

	if s.IsDisjunction {
		fmt.Fprintf(&b, `
		
			return json.Marshal(json_%s)
		}

		`, s.Name)
	} else {
		fmt.Fprintf(&b, `
	
			model := &json%s {
		%s
			}
			return json.Marshal(model)
		}
	
		`, s.Name, strings.Join(structLines, ""))
	}

	return b.String()
}

func (s *Model) generateGetAttrFunction(node Node) string {
	b := strings.Builder{}
	attrName := utils.ToCamelCase(node.Name)

	fmt.Fprintf(&b, "func (m %s) GetAttr%s() interface{} {\n", s.Name, attrName)
	b.WriteString("var attr interface{}\nvar err error\n\n")

	for _, kind := range node.DisjunctionKinds {
		switch kind {
		case cue.StructKind, cue.ListKind:
			fmt.Fprintf(&b, "err = json.Unmarshal([]byte(m.%s.ValueString()), &attr)", attrName)
		case cue.BoolKind:
			fmt.Fprintf(&b, "attr, err = strconv.ParseBool(m.%s.ValueString())", attrName)
		case cue.IntKind:
			fmt.Fprintf(&b, "attr, err = strconv.ParseInt(m.%s.ValueString(), 10, 64)", attrName)
		case cue.NumberKind, cue.FloatKind:
			fmt.Fprintf(&b, "attr, err = strconv.ParseFloat(m.%s.ValueString(), 64)", attrName)
		case cue.StringKind:
			continue
		}

		b.WriteString(`
			if err == nil {
				return attr
			}
		`)
	}
	fmt.Fprintf(&b, `
			return m.%s.ValueString()
		}

	`, attrName)

	return b.String()
}

func (s *Model) generateDefaultsFunction() string {
	defaults := make([]string, 0)
	for _, node := range s.Nodes {
		kind := node.TerraformType()

		if kind != "" && node.Default != "" {
			defaults = append(defaults, fmt.Sprintf(`if m.%s.IsNull() {
	m.%s = types.%sValue(%s)
}`, utils.ToCamelCase(node.Name), utils.ToCamelCase(node.Name), kind, node.Default))
		}

		if node.Kind == cue.ListKind && node.SubTerraformType() != "" {
			defaults = append(defaults, fmt.Sprintf(`if len(m.%s.Elements()) == 0 {
	m.%s, _ = types.ListValue(types.%sType, []attr.Value{})
}`, utils.ToCamelCase(node.Name), utils.ToCamelCase(node.Name), node.SubTerraformType()))
		}

	}

	return fmt.Sprintf(`func (m %[1]s) ApplyDefaults() %[1]s {
	%s
	return m
}

`, s.Name, strings.Join(defaults, "\n"))
}

func (s *Model) Generate() string {
	b := strings.Builder{}
	for _, node := range s.Nodes {
		if node.Kind == cue.StructKind || node.Kind == cue.ListKind && node.SubKind == cue.StructKind {
			nestedModel := Model{
				Name:          s.Name + "_" + utils.Title(node.Name),
				IsDisjunction: node.IsDisjunction,
				Nodes:         node.Children,
				Nested:        true,
			}
			b.WriteString(nestedModel.Generate())
		}
	}

	b.WriteString(s.terraformModel())
	b.WriteString(s.generateToJSONFunction())
	b.WriteString(s.generateDefaultsFunction())

	return b.String()
}
