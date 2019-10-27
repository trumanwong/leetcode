package main

import (
	"leetcode/algorithms/0080.RemoveDuplicatesfromSortedArrayII/removeDuplicates"
	"testing"
)

func TestRemoveDuplicates(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		ret := removeDuplicates.RemoveDuplicates(test.nums)
		if ret != test.output {
			if ret != test.output {
				t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
			}
		}
	}
}