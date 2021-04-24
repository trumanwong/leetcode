package main

import (
	"leetcode/algorithms/0377.CombinationSumIV/combinationSum4"
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
	}

	for _, test := range tests {
		ret := combinationSum4.CombinationSum4(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.nums, test.target, test.output)
		}
	}
}
