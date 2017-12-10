package main

import (
	"fmt"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestEscape(t *testing.T) {
	for i, c := range []struct {
		input    []int
		expected int64
	}{
		{input: []int{0}, expected: 2},
		{input: []int{0, 3, 0, 1, -3}, expected: 5},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			assert.Equal(t, escapeTheMaze(c.input), c.expected)
		})
	}
}
