package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	config := GeneratorConfig{
		Suffix: "Event",
	}

	var flags flag.FlagSet
	flags.StringVar(&config.Suffix, "suffixMatch", config.Suffix, "Suffix required for a message to be handled as an event. Set to empty to generate for all messages.")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generateFile(gen, f, config)
		}
		return nil
	})
}
