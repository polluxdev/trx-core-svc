package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CamelCase", "camel_case"},
		{"camelCase", "camel_case"},
		{"CamelCaseWithNumbers123", "camel_case_with_numbers123"},
		{"", ""},
		{"Already_Snake_Case", "already_snake_case"},
	}

	for _, test := range tests {
		output := ToSnakeCase(test.input)
		assert.Equal(t, test.expected, output, "Expected %s but got %s", test.expected, output)
	}
}

func TestFormatDateTime(t *testing.T) {
	format := "2006-01-02 15:04:05"

	tests := []struct {
		input    *time.Time
		expected *string
	}{
		{
			input:    func() *time.Time { t := time.Date(2024, 7, 1, 10, 0, 0, 0, time.UTC); return &t }(),
			expected: func() *string { s := "2024-07-01 10:00:00"; return &s }(),
		},
		{
			input:    nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		output := FormatDateTime(test.input, format)
		assert.Equal(t, test.expected, output, "Expected %s but got %s", test.expected, output)
	}
}
