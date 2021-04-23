package main

import (
	"leetcode/algorithms/1133.LargestUniqueNumber/largestUniqueNumber"
	"testing"
)

func TestLargestUniqueNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{5, 7, 3, 9, 4, 9, 8, 3, 1}, 8},
		{[]int{9, 9, 8, 8}, -1},
	}

	for _, test := range tests {
		ret := largestUniqueNumber.LargestUniqueNumber(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
