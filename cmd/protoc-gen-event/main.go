package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	config := GeneratorConfig{
		Suffix:         "Event",
		RequiredFields: []requiredField{},
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
