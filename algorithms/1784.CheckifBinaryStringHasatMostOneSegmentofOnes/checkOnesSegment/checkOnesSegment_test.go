package checkOnesSegment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckOnesSegment(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{
		{"1001", false},
		{"110", true},
		{"1", true},
		{"10", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, checkOnesSegment(test.s))
	}
}
