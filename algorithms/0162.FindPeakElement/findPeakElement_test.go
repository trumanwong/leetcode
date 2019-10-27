package main

import (
	"leetcode/algorithms/0162.FindPeakElement/findPeakElement"
	"testing"
)

func TestFindPeakElement(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}

	for _, test := range tests {
		ret := findPeakElement.FindPeakElement(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
