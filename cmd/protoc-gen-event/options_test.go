package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRequiredFieldsFlag(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected requiredFields
	}{
		{
			name:     "no input",
			input:    "",
			expected: nil,
		},
		{
			name:  "single field",
			input: "string:messageID",
			expected: requiredFields{
				{
					Name: "messageID",
					Type: "string",
				},
			},
		},
		{
			name:  "multiple fields",
			input: "string:messageID+google.protobuf.Timestamp:generatedAt",
			expected: requiredFields{
				{
					Name: "messageID",
					Type: "string",
				},
				{
					Name: "generatedAt",
					Type: "google.protobuf.Timestamp",
				},
			},
		},

		{
			name:     "invalid single field",
			input:    "string&messageID",
			expected: requiredFields{},
		},
		{
			name:     "missing field name",
			input:    "string:",
			expected: requiredFields{},
		},
		{
			name:     "too much data",
			input:    "string:fieldName:what_is_that",
			expected: requiredFields{},
		},
		{
			name:  "invalid field with valid field",
			input: "string:+google.protobuf.Timestamp:generatedAt",
			expected: requiredFields{
				{
					Name: "generatedAt",
					Type: "google.protobuf.Timestamp",
				},
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			result := parseRequiredFieldsFlag(tc.input)

			require.Equal(t, tc.expected, result)
		})
	}
}
