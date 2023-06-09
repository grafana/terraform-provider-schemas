//go:generate go run gen.go

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grafana/codejen"
	"github.com/grafana/grafana/pkg/plugins/pfs/corelist"
	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/kindsys"
	jen "github.com/grafana/terraform-provider-schemas/gen/internal"
	"github.com/grafana/terraform-provider-schemas/gen/terraform"
)

func main() {
	// Load all core and composable kinds.
	var corek []kindsys.Kind
	var compok []kindsys.Composable

	minMaturity :=
		func() kindsys.Maturity {
			switch os.Getenv("MIN_MATURITY") {
			case "merged":
				return kindsys.MaturityMerged
			case "experimental":
				return kindsys.MaturityExperimental
			case "stable":
				return kindsys.MaturityStable
			case "mature":
				return kindsys.MaturityMature
			default:
				return kindsys.MaturityExperimental
			}
		}()

	for _, kind := range corekind.NewBase(nil).All() {
		// This provider should only generate dashboard-related datasources for now
		if kind.Name() != "Dashboard" {
			continue
		}
		if kind.Maturity().Less(minMaturity) {
			continue
		}
		corek = append(corek, kind)
	}

	// Skip data sources that only have definitions
	skippedPlugins := map[string]bool{
		"AzureMonitorDataQuery":          true,
		"CloudWatchDataQuery":            true,
		"GoogleCloudMonitoringDataQuery": true,
		"TempoDataQuery":                 true,
	}
	for _, pp := range corelist.New(nil) {
		for _, kind := range pp.ComposableKinds {
			_, skipped := skippedPlugins[kind.Lineage().Name()]
			if kind.Maturity().Less(minMaturity) || skipped {
				continue
			}
			compok = append(compok, kind)
		}
	}

	// Add all jennies.
	jfs := codejen.NewFS()
	tj := getJennies()

	ckfs, err := tj.Core.GenerateFS(corek...)
	die(err)
	die(jfs.Merge(ckfs))
	ckfs, err = tj.Composable.GenerateFS(compok...)
	die(err)
	die(jfs.Merge(ckfs))

	if _, set := os.LookupEnv("CODEGEN_VERIFY"); set {
		if err = jfs.Verify(context.Background(), ""); err != nil {
			die(fmt.Errorf("generated code is out of sync with inputs:\n%s\nrun `make gen-cue` to regenerate", err))
		}
	} else if err = jfs.Write(context.Background(), ""); err != nil {
		die(fmt.Errorf("error while writing generated code to disk:\n%s", err))
	}
}

// Get all jennies for Terraform.
func getJennies() jen.TargetJennies {
	header := "terraform"
	path := "../internal/provider"
	tj := terraform.JenniesForTerraform()

	tj.Core.AddPostprocessors(jen.Prefixer(path), jen.SlashHeaderMapper(header))
	tj.Composable.AddPostprocessors(jen.Prefixer(path), jen.SlashHeaderMapper(header))

	return tj
}

func die(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}
}
