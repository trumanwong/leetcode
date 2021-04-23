package main

import (
	"leetcode/algorithms/0704.BinarySearch/search"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, test := range tests {
		ret := search.Search(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for nums = %v, target = %d; expected %d", ret, test.nums, test.target, test.output)
		}
	}
}
