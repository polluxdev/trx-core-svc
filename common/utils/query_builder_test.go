package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructConditionalClause(t *testing.T) {
	tests := []struct {
		name     string
		builder  []*ConditionalBuilder
		expected string
		args     []interface{}
	}{
		{
			name: "LIKE clause",
			builder: []*ConditionalBuilder{
				{Column: "name", Value: "Windah", Logical: "LIKE", Operator: "AND"},
			},
			expected: "1 = 1 AND name LIKE ?",
			args:     []interface{}{"%Windah%"},
		},
		{
			name: "IN clause",
			builder: []*ConditionalBuilder{
				{Column: "id", Value: []int{1, 2, 3}, Logical: "IN", Operator: "AND"},
			},
			expected: "1 = 1 AND id IN (?)",
			args:     []interface{}{[]int{1, 2, 3}},
		},
		{
			name: "BETWEEN clause",
			builder: []*ConditionalBuilder{
				{Column: "date", Value: []string{"2023-01-01", "2023-12-31"}, Logical: "BETWEEN", Operator: "AND"},
			},
			expected: "1 = 1 AND date BETWEEN ? AND ?",
			args:     []interface{}{[]string{"2023-01-01", "2023-12-31"}},
		},
		{
			name: "IS NULL clause",
			builder: []*ConditionalBuilder{
				{Column: "deleted_at", Logical: "IS NULL", Operator: "AND"},
			},
			expected: "1 = 1 AND deleted_at IS NULL",
			args:     []interface{}{},
		},
		{
			name: "IS NOT NULL clause",
			builder: []*ConditionalBuilder{
				{Column: "deleted_at", Logical: "IS NOT NULL", Operator: "AND"},
			},
			expected: "1 = 1 AND deleted_at IS NOT NULL",
			args:     []interface{}{},
		},
		{
			name: "Default clause",
			builder: []*ConditionalBuilder{
				{Column: "age", Value: 30, Logical: "=", Operator: "AND"},
			},
			expected: "1 = 1 AND age = ?",
			args:     []interface{}{30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, args := ConstructConditionalClause(tt.builder)
			assert.Equal(t, tt.expected, result)
			if args == nil {
				args = []interface{}{}
			}
			assert.Equal(t, tt.args, args)
		})
	}
}

func TestSetDefaultClause(t *testing.T) {
	expected := "1 = 1"
	result := SetDefaultClause()
	assert.Equal(t, expected, result)
}
