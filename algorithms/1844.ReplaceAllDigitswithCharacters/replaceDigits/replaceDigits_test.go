package replaceDigits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceDigits(t *testing.T)  {
	tests := []struct{
		s string
		output string
	}{
		{"a1c1e1", "abcdef"},
		{"a1b2c3d4e", "abbdcfdhe"},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, replaceDigits(test.s))
	}
}