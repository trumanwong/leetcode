package main

import (
	"testing"
	"truman.com/leetcode/503.NextGreaterElementII/nextGreaterElements"
)

func TestNextGreaterElements(t *testing.T)  {
	tests := []struct{
		nums []int
		output []int
	}{
		{[]int{1,2,1},[]int{2,-1,2}},
	}

	for _, test := range tests {
		ret := nextGreaterElements.NextGreaterElements(test.nums)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
				break
			}
		}
	}
}