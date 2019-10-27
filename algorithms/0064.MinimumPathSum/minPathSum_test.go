package main

import (
	"leetcode/algorithms/0064.MinimumPathSum/minPathSum"
	"testing"
)

func TestMinPathSum(t *testing.T)  {
	tests := []struct{
		input [][]int
		output int
	} {
		{[][]int{
			{1,3,1},
			{1,5,1},
			{4,2,1},
		}, 7},
	}

	for _, test := range tests {
		ret := minPathSum.MinPathSum(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}