package main

import (
	"leetcode/algorithms/0962.MaximumWidthRamp/maxWidthRamp"
	"testing"
)

func TestMaxWidthRamp(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{6,0,8,2,1,5},4},
		{[]int{9,8,1,0,1,9,4,0,4,1}, 7},
	}
	for _, test := range tests {
		ret := maxWidthRamp.MaxWidthRamp(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}