package main

import (
	"leetcode/algorithms/0747.LargestNumberAtLeastTwiceofOthers/dominantIndex"
	"testing"
)

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}

	for _, test := range tests {
		ret := dominantIndex.DominantIndex(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
