package main

import (
	"fmt"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestRowChecksum(t *testing.T) {
	for _, c := range []struct {
		input    []int64
		expected int64
	}{
		{input: []int64{}, expected: 0},
		{input: []int64{5, 1, 9, 5}, expected: 8},
		{input: []int64{7, 5, 3}, expected: 4},
		{input: []int64{2, 4, 6, 8}, expected: 6},
	} {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			assert.Equal(t, rowChecksum(c.input), c.expected)
		})
	}
}
