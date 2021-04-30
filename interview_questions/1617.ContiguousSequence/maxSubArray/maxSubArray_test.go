package maxSubArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, maxSubArray(test.nums))
	}
}
