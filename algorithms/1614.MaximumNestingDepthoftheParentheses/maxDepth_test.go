package main

import (
	"leetcode/algorithms/1614.MaximumNestingDepthoftheParentheses/maxDepth"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"1+(2*3)/(2-1)", 1},
		{"1", 0},
	}

	for _, test := range tests {
		ret := maxDepth.MaxDepth(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}
