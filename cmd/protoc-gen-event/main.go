package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	config := GeneratorConfig{
		Suffix:         "Event",
		RequiredFields: []RequiredField{},
	}

	var flags flag.FlagSet
	var requiredFields string
	flags.StringVar(&config.Suffix, "suffixMatch", config.Suffix, "Suffix required for a message to be handled as an event. Set to empty to generate for all messages.")
	flags.StringVar(&requiredFields, "requiredFields", "", "Mandatory fields in events. Example: 'SomeFieldName@string+EmittedAt@google.protobuf.Timestamp'")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		config.RequiredFields = parseRequiredFieldsFlag(requiredFields)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			_, err := generateFile(gen, f, config)
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error()+"\n")
				os.Exit(1)
			}
		}

		return nil
	})
}

func parseRequiredFieldsFlag(flagInput string) []RequiredField {
	if flagInput == "" {
		return nil
	}

	fields := strings.Split(flagInput, "+")
	requiredFields := make([]RequiredField, 0, len(fields))

	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			continue
		}

		requiredFields = append(requiredFields, RequiredField{
			Type: parts[0],
			Name: parts[1],
		})
	}

	return requiredFields
}
