package mergeAlternately

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		excepted string
	}{
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, mergeAlternately(test.word1, test.word2))
	}
}
