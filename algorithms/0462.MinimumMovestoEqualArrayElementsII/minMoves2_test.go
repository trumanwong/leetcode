package main

import (
	"leetcode/algorithms/0462.MinimumMovestoEqualArrayElementsII/minMoves2"
	"testing"
)

func TestMinMoves2(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 2, 3}, 2},
	}

	for _, test := range tests {
		ret := minMoves2.MinMoves2(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret ,test.nums, test.output)
		}
	}
}