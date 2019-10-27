package main

import (
	"leetcode/algorithms/0153.FindMinimuminRotatedSortedArray/findMin"
	"testing"
)

func TestFindMin(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}

	for _, test := range tests {
		ret := findMin.FindMin(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}