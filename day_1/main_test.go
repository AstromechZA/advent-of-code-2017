package main

import (
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestCaptcha(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected int
	}{
		{input: "", expected: 0},
		{input: "1122", expected: 3},
		{input: "1111", expected: 4},
		{input: "1234", expected: 0},
		{input: "91212129", expected: 9},
	} {
		t.Run(c.input, func(t *testing.T) {
			assert.Equal(t, generateCaptcha(c.input), c.expected)
		})
	}
}
