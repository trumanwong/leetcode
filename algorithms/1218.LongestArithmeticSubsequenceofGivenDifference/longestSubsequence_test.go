package main

import (
	"leetcode/algorithms/1218.LongestArithmeticSubsequenceofGivenDifference/longestSubsequence"
	"testing"
)

func TestLongestSubsequence(t *testing.T)  {
	tests := []struct{
		arr []int
		difference int
		output int
	}{
		{[]int{1, 2, 3, 4}, 1, 4},
		{[]int{1, 3, 5, 7}, 1, 1},
		{[]int{1,5,7,8,5,3,4,2,1}, -2, 4},
	}

	for _, test := range tests {
		ret := longestSubsequence.LongestSubsequence(test.arr, test.difference)
		if ret != test.output {
			t.Errorf("Got %d for input arr = %v, diffrence = %d; expected %d", ret, test.arr, test.difference, test.output)
		}
	}
}