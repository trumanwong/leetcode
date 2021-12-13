package strangePrinter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrangePrinter(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"aaabbb", 2},
		{"aba", 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, strangePrinter(test.s))
	}
}
