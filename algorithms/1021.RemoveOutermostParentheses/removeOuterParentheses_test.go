package main

import (
	"leetcode/algorithms/1021.RemoveOutermostParentheses/removeOuterParentheses"
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
		{"((()())(()()))", "(()())(()())"},
	}

	for _, test := range tests {
		ret := removeOuterParentheses.RemoveOuterParentheses(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}
