package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNow(t *testing.T) {
	now := time.Now()
	result := GetNow(0)
	assert.WithinDuration(t, now, *result, time.Second, "Expected time to be within one second of now")

	duration := 5 * time.Minute
	expected := now.Add(duration)
	result = GetNow(duration)
	assert.WithinDuration(t, expected, *result, time.Second, "Expected time to be within one second of now plus duration")

	duration = -5 * time.Minute
	expected = now.Add(duration)
	result = GetNow(duration)
	assert.WithinDuration(t, expected, *result, time.Second, "Expected time to be within one second of now minus duration")
}
