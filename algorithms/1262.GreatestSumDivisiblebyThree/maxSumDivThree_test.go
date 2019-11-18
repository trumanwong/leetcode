package main

import (
	"leetcode/algorithms/1262.GreatestSumDivisiblebyThree/maxSumDivThree"
	"testing"
)

func TestMaxSumDivThree(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		//{[]int{3, 6, 5, 8, 1}, 18},
		{[]int{3, 6, 5, 1, 8}, 18},
	}

	for _, test := range tests {
		ret := maxSumDivThree.MaxSumDivThree(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}