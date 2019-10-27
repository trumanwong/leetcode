package main

import (
	"leetcode/algorithms/0016.3SumCloses/threeSumClosest"
	"testing"
)

func TestThreeSumClosest(t *testing.T)  {
	tests := []struct{
		nums []int
		target int
		output int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
	}

	for _, test := range tests {
		ret := threeSumClosest.ThreeSumClosest(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input %v, target = %d; expected %d", ret, test.nums, test.target, test.output)
		}
	}
}