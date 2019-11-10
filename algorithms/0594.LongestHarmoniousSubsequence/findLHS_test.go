package main

import (
	"leetcode/algorithms/0594.LongestHarmoniousSubsequence/findLHS"
	"testing"
)

func TestFindLHS(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1,3,2,2,5,2,3,7}, 5},
	}

	for _, test := range tests {
		ret := findLHS.FindLHS(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}