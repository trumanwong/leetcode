package main

import (
	"leetcode/algorithms/0566.ReshapetheMatrix/matrixReshape"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		nums   [][]int
		r      int
		c      int
		output [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
	}

	for _, test := range tests {
		ret := matrixReshape.MatrixReshape(test.nums, test.r, test.c)
		for i := 0; i < len(ret); i++ {
			for j := 0; j < len(ret[i]); j++ {
				if ret[i][j] != test.output[i][j] {
					t.Errorf("Got %v for input %v, r = %d, c = %d; expected %v", ret, test.nums, test.r, test.c, test.output)
					break
				}
			}
		}
	}
}
