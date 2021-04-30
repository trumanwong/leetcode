package minCount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCount(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{4, 2, 1}, 4},
		{[]int{2, 3, 10}, 8},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, minCount(test.nums))
	}
}
