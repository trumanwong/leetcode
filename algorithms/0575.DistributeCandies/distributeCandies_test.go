package main

import (
	"leetcode/algorithms/0575.DistributeCandies/distributeCandies"
	"testing"
)

func TestDistributeCandies(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
	}

	for _, test := range tests {
		ret := distributeCandies.DistributeCandies(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}