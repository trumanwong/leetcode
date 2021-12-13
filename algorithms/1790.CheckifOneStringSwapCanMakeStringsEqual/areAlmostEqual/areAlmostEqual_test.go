package areAlmostEqual

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreAlmostEqual(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		excepted bool
	}{
		{"bank", "kanb", true},
		{"attack", "defend", false},
		{"kelb", "kelb", true},
		{"abcd", "dcba", false},
		{"aa", "ac", false},
		{"ab", "ba", true},
	}
	for _, test := range tests {
		assert.Equal(t, test.excepted, areAlmostEqual(test.s1, test.s2), test.s1 + ":" + test.s2)
	}
}
