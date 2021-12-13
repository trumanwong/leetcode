package maxAscendingSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxAscendingSum(t *testing.T) {
	tests := []struct {
		nums     []int
		excepted int
	}{
		{[]int{10,20,30,5,10,50}, 65},
		{[]int{10,20,30,40,50}, 150},
		{[]int{12,17,15,13,10,11,12}, 33},
		{[]int{100,10,1}, 100},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, maxAscendingSum(test.nums))
	}
}
