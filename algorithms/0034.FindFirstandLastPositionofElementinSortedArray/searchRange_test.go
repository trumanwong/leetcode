package main

import (
	"leetcode/algorithms/0034.FindFirstandLastPositionofElementinSortedArray/searchRange"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}

	for _, test := range tests {
		ret := searchRange.SearchRange(test.nums, test.target)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v, target=%d;expected %v", ret, test.nums, test.target, test.output)
				break
			}
		}
	}
}
