package main

import (
	"strings"
)

func parseRequiredFieldsFlag(flagInput string) RequiredFields {
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

		if len(parts[0]) == 0 || len(parts[1]) == 0 {
			continue
		}

		requiredFields = append(requiredFields, RequiredField{
			Type: parts[0],
			Name: parts[1],
		})
	}

	return requiredFields
}
