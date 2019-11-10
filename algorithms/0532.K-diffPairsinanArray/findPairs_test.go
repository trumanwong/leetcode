package main

import (
	"leetcode/algorithms/0532.K-diffPairsinanArray/findPairs"
	"testing"
)

func TestFindPairs(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		output int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
	}

	for _, test := range tests {
		ret := findPairs.FindPairs(test.nums, test.k)
		if ret != test.output {
			t.Errorf("Got %d for input %v, k = %d; expected %d", ret, test.nums, test.k, test.output)
		}
	}
}