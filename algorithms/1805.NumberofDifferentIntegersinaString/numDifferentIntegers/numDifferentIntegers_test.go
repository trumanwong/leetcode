package numDifferentIntegers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDifferentIntegers(t *testing.T) {
	tests := []struct {
		word     string
		excepted int
	}{
		{"a123bc34d8ef34", 3},
		{"leet1234code234", 2},
		{"a1b01c001", 1},
		{"192383183928778851682383a2089984061937879119", 2},
		{"00000a", 1},
		{"gi851a851q8510v", 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, numDifferentIntegers(test.word))
	}
}
