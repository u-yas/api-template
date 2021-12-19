package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/iancoleman/strcase"
)

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {

	for _, model := range b.Models {
		for _, field := range model.Fields {
			str := field.Tag
			regJson := regexp.MustCompile("json:")
			regDoubleQuotes := regexp.MustCompile("\"")

			removedJsonTag := regJson.ReplaceAllString(str, "")
			removedDoubleQuotes := regDoubleQuotes.ReplaceAllString(removedJsonTag, "")
			snack_case_tag := strcase.ToSnake(removedDoubleQuotes)

			fmt.Println(snack_case_tag)
			field.Tag = "json:" + "\"" + snack_case_tag + "\""
		}
	}

	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
