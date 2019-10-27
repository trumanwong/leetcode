package main

import (
	"leetcode/algorithms/0027.RemoveElement/removeElement"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct{
		nums []int
		val int
		output int
	} {
		{[]int{3,2,2,3}, 3, 2},
		{[]int{0,1,2,2,3,0,4,2}, 2, 5},
	}

	for _, test := range tests {
		ret := removeElement.RemoveElement(test.nums, test.val)

		if ret != test.output {
			t.Errorf("Got %d for input %v, val = %d; expected %d", ret, test.nums, test.val, test.output)
		}
	}
}