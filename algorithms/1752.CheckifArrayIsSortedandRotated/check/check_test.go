package check

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		nums     []int
		excepted bool
	}{
		{[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 1, 1}, true},
		{[]int{2, 1}, true},
		{[]int{3,2,1}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, check(test.nums))
	}
}
