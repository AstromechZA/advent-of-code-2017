package main

import (
	"fmt"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestDataDistance(t *testing.T) {
	for _, c := range []struct {
		input    int64
		expected int64
	}{
		{input: 1, expected: 0},
		{input: 12, expected: 3},
		{input: 23, expected: 2},
		{input: 1024, expected: 31},
	} {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			assert.Equal(t, DataDistance(c.input), c.expected)
		})
	}
}
