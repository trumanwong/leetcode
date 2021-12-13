package countConsistentStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountConsistentStrings(t *testing.T) {
	tests := []struct {
		allowed  string
		words    []string
		excepted int
	}{
		{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2},
		{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7},
		{"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, countConsistentStrings(test.allowed, test.words))
	}
}
