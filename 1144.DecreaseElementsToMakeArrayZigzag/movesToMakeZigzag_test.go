package main

import (
	"testing"
	"truman.com/leetcode/1144.DecreaseElementsToMakeArrayZigzag/movesToMakeZigzag"
)

func TestTribonacci(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1,2,3}, 2},
		{[]int{2,7,10,9,8,9}, 4},
	}

	for _, test := range tests {
		ret := movesToMakeZigzag.MovesToMakeZigzag(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}

	}
}