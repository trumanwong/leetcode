package main

import (
	"testing"
	"truman.com/leetcode/1021.RemoveOutermostParentheses/removeOuterParentheses"
)

func TestRemoveOuterParentheses(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"(()())(())","()()()"},
		{"(()())(())(()(()))","()()()()(())"},
		{"()()",""},
		{"((()())(()()))","(()())(()())"},
	}

	for _, test := range tests {
		ret := removeOuterParentheses.RemoveOuterParentheses(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}