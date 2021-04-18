package main

import (
	"leetcode/algorithms/1608.SpecialArrayWithXElementsGreaterThanorEqualX/specialArray"
	"testing"
)

func TestSpecialArray(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{3, 5}, 2},
		{[]int{0, 0}, -1},
		{[]int{0, 4, 3, 0, 4}, 3},
		{[]int{3, 6, 7, 7, 0}, -1},
	}

	for _, test := range tests {
		ret := specialArray.SpecialArray(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}