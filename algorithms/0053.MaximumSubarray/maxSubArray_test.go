package main

import (
	"leetcode/algorithms/0053.MaximumSubarray/maxSubArray"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, test := range tests {
		ret := maxSubArray.MaxSubArray(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
