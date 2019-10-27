package main

import (
	"leetcode/algorithms/1150.CheckIfaNumberIsMajorityElementinaSortedArray/isMajorityElement"
	"testing"
)

func TestIsMajorityElement(t *testing.T)  {
	tests := []struct{
		nums []int
		target int
		output bool
	}{
		{[]int{2,4,5,5,5,5,5,6,6}, 5, true},
		{[]int{10,100,101,101}, 101, false},
	}

	for _, test := range tests {
		ret := isMajorityElement.IsMajorityElement(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %t for input nums = %v, target = %d; expected %t", ret, test.nums, test.target, test.output)
		}
	}
}