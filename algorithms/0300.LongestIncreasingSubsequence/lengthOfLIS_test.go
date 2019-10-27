package main

import (
	"leetcode/algorithms/0300.LongestIncreasingSubsequence/lengthOfLIS"
	"testing"
)

func TestLengthOfLIS(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}

	for _, test := range tests {
		ret := lengthOfLIS.LengthOfLIS(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}