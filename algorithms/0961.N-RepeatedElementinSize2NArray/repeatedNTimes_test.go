package main

import (
	"leetcode/algorithms/0961.N-RepeatedElementinSize2NArray/repeatedNTimes"
	"testing"
)

func TestRepeatedNTimes(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{1,2,3,3},3},
		{[]int{2,1,2,5,3,2},2},
		{[]int{5,1,5,2,5,3,5,4},5},
	}
	for _, test := range tests {
		ret := repeatedNTimes.RepeatedNTimes(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}