package main

import (
	"leetcode/algorithms/0581.ShortestUnsortedContinuousSubarray/findUnsortedSubarray"
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
	}

	for _, test := range tests {
		ret := findUnsortedSubarray.FindUnsortedSubarray(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}