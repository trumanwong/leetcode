package secondHighest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecondHighest(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"dfa12321afd", 2},
		{"abc1111", -1},
		{"sjhtz8344", 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, secondHighest(test.s))
	}
}
