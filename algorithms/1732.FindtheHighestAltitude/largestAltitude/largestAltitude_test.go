package largestAltitude

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain     []int
		excepted int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}
	for _, test := range tests {
		assert.Equal(t, test.excepted, largestAltitude(test.gain))
	}
}
