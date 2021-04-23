package main

import (
	"leetcode/algorithms/1111.MaximumNestingDepthofTwoValidParenthesesStrings/maxDepthAfterSplit"
	"testing"
)

func TestMaxDepthAfterSplit(t *testing.T) {
	tests := []struct {
		seq    string
		output []int
	}{
		{"(()())", []int{0, 1, 1, 1, 1, 0}},
		{"()(())()", []int{0, 0, 0, 1, 1, 0, 0, 0}},
	}

	for _, test := range tests {
		ret := maxDepthAfterSplit.MaxDepthAfterSplit(test.seq)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s; expected %v", ret, test.seq, test.output)
				break
			}
		}
	}
}
