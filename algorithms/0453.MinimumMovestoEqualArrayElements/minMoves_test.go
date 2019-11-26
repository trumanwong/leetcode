package main

import (
	"leetcode/algorithms/0453.MinimumMovestoEqualArrayElements/minMoves"
	"testing"
)

func TestMinMoves(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{3, 2, 1}, 3},
	}

	for _, test := range tests {
		ret := minMoves.MinMoves(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}