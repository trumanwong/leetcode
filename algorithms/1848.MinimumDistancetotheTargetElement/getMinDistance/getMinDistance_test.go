package getMinDistance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinDistance(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		start  int
		output int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 3, 1},
		{[]int{1}, 1, 0, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, getMinDistance(test.nums, test.target, test.start))
	}
}
