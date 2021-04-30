package arraySign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySign(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{-1,-2,-3,-4,3,2,1},1},
		{[]int{1,5,0,2,-3}, 0},
		{[]int{-1,1,-1,1,-1}, -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, arraySign(test.nums))
	}
}
