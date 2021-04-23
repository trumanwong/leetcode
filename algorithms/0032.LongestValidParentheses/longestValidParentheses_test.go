package main

import (
	"leetcode/algorithms/0032.LongestValidParentheses/longestValidParentheses"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"(()", 2},
		{")()())", 4},
	}

	for _, test := range tests {
		ret := longestValidParentheses.LongestValidParentheses(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}
