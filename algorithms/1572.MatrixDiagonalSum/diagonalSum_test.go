package main

import (
	"leetcode/algorithms/1572.MatrixDiagonalSum/diagonalSum"
	"testing"
)

func TestDiagonalSum(t *testing.T) {
	tests := []struct {
		mat    [][]int
		output int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 25},
		{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, 8},
		{[][]int{{5}}, 5},
	}

	for _, test := range tests {
		ret := diagonalSum.DiagonalSum(test.mat)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.mat, test.output)
		}
	}
}
