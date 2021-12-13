package maxRepeating

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxRepeating(t *testing.T) {
	tests := []struct {
		sequence string
		word     string
		excepted int
	}{
		{"ababc", "ab", 2},
		{"ababc", "ba", 1},
		{"ababc", "ac", 0},
		{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba", 5},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, maxRepeating(test.sequence, test.word))
	}
}
