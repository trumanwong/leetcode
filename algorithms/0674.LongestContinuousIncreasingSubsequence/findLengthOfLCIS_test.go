package main

import (
	"leetcode/algorithms/0674.LongestContinuousIncreasingSubsequence/findLengthOfLCIS"
	"testing"
)

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
	}

	for _, test := range tests {
		ret := findLengthOfLCIS.FindLengthOfLCIS(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
