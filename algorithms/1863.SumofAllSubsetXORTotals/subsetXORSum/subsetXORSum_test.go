package subsetXORSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetXORSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3}, 6},
		{[]int{5, 1, 6}, 28},
		{[]int{3, 4, 5, 6, 7, 8}, 480},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, subsetXORSum(test.nums))
	}
}
