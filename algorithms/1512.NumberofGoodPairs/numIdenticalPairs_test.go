package main

import (
	"leetcode/algorithms/1512.NumberofGoodPairs/numIdenticalPairs"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, test := range tests {
		ret := numIdenticalPairs.NumIdenticalPairs(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
