package main

import (
	"testing"
	"truman.com/leetcode/287.FindtheDuplicateNumber/findDuplicate"
)

func TestFindDuplicate(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	} {
		{[]int{1,3,4,2,2}, 2},
		{[]int{3,1,3,4,2}, 3},
	}

	for _, test := range tests {
		ret := findDuplicate.FindDuplicate(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}