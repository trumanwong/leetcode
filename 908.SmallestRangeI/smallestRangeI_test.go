package main

import (
	"testing"
	"truman.com/leetcode/908.SmallestRangeI/smallestRangeI"
)

func TestSmallestRangeI(t *testing.T)  {
	tests := []struct{
		A []int
		K int
		output int
	} {
		{[]int{1}, 0, 0},
		{[]int{0, 10}, 2, 6},
		{[]int{1,3,6}, 3,0},
		{[]int{2,7,2}, 1, 3},
		{[]int{3,1,10}, 4,1},
		{[]int{7,8,8}, 5,0},
	}

	for _, test := range tests {
		ret := smallestRangeI.SmallestRangeI(test.A, test.K)
		if ret != test.output {
			t.Errorf("Got %d for input A = %v, K = %d; expected %v", ret, test.A, test.K, test.output)
		}
	}
}