package xorQueries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXorQueries(t *testing.T) {
	tests := []struct {
		arr      []int
		queries  [][]int
		excepted []int
	}{
		{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}, []int{2, 7, 14, 8}},
		{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}, []int{8, 0, 4, 4}},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, xorQueries(test.arr, test.queries))
	}
}
