package xorOperation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXorOperation(t *testing.T) {
	tests := []struct {
		n        int
		start    int
		excepted int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, xorOperation(test.n, test.start))
	}
}
