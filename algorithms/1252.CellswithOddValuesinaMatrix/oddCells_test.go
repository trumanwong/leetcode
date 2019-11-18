package main

import (
	"leetcode/algorithms/1252.CellswithOddValuesinaMatrix/oddCells"
	"testing"
)

func TestOddCells(t *testing.T)  {
	tests := []struct{
		n int
		m int
		indices [][]int
		output int
	}{
		{2, 3, [][]int{{0,1},{1,1}}, 6},
		{2,2,[][]int{{1,1},{0,0}}, 0},
	}

	for _, test := range tests {
		ret := oddCells.OddCells(test.n, test.m, test.indices)
		if ret != test.output {
			t.Errorf("Got %d for input m = %d, n = %d, indices = %v; expected %d", ret, test.m, test.n, test.indices, test.output)
		}
	}
}