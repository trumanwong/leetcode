package main

import (
	"leetcode/algorithms/0724.FindPivotIndex/pivotIndex"
	"testing"
)

func TestPivotIndex(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1,2,3}, -1},
	}

	for _, test := range tests {
		ret := pivotIndex.PivotIndex(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}