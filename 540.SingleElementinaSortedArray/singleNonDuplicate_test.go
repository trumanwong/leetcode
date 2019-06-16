package main

import (
	"testing"
	"truman.com/leetcode/540.SingleElementinaSortedArray/singleNonDuplicate"
)

func TestSingleNonDuplicate(t *testing.T)  {
	tests := []struct {
		Input  []int
		Output int
	}{
		{[]int{1,1,2,3,3,4,4,8,8}, 2},
		{[]int{3,3,7,7,10,11,11}, 10},
	}

	for _, test := range tests {
		ret := singleNonDuplicate.SingleNonDuplicate(test.Input)
		if ret != test.Output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.Input, test.Output)
		}
	}
}