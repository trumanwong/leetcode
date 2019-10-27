package main

import (
	"leetcode/algorithms/1027.LongestArithmeticSequence/longestArithSeqLength"
	"testing"
)

func TestLongestArithSeqLength(t *testing.T)  {
	tests := []struct {
		input []int
		output int
	}{
		{[]int{3,6,9,12}, 4},
		{[]int{9,4,7,2,10},3},
		{[]int{20,1,15,3,10,5,8},4},
	}

	for _, test := range tests {
		ret := longestArithSeqLength.LongestArithSeqLength(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}