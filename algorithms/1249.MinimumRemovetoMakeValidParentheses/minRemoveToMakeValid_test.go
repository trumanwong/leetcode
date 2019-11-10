package main

import (
	"leetcode/algorithms/1249.MinimumRemovetoMakeValidParentheses/minRemoveToMakeValid"
	"testing"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct{
		s string
		output string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
		{"(a(b(c)d)", "(a(bc)d)"},
		{"())()(((", "()()"},
	}

	for _, test := range tests {
		ret := minRemoveToMakeValid.MinRemoveToMakeValid(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}