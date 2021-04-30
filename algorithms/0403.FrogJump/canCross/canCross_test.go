package canCross

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCross(t *testing.T) {
	tests := []struct {
		stones []int
		output bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, canCross(test.stones))
	}
}
