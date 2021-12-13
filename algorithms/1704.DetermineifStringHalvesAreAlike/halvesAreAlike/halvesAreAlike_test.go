package halvesAreAlike

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHalvesAreAlike(t *testing.T) {
	tests := []struct {
		s        string
		excepted bool
	}{
		{"book", true},
		{"textbook", false},
		{"MerryChristmas", false},
		{"AbCdEfGh", true},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, halvesAreAlike(test.s))
	}
}
