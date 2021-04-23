package main

import (
	"leetcode/algorithms/0022.GenerateParentheses/generateParenthesis"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n      int
		output []string
	}{
		{3, []string{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"}},
	}

	for _, test := range tests {
		ret := generateParenthesis.GenerateParenthesis(test.n)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %d; expected %v", ret, test.n, test.output)
			}
		}
	}
}
