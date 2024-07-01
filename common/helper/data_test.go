package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaultIfZero(t *testing.T) {
	tests := []struct {
		name         string
		value        int
		defaultValue int
		expected     int
	}{
		{"zero value", 0, 0, 0},
		{"non zero value", 1, 0, 1},
		{"non zero value", -1, 1, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SetDefaultIfZero(test.value, test.defaultValue)
			assert.Equal(t, test.expected, result)
		})
	}
}
