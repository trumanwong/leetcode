package main

import (
	"testing"
	"truman.com/leetcode/20.ValidParentheses/isValid"
)

func TestIsValide(t *testing.T)  {
	tests := []struct{
		input string
		output bool
	}{
		{"()",true},
		{"()[]{}",true},
		{"(]",false},
		{"([)]",false},
		{"{[]}",true},
	}

	for _, test := range tests {
		ret := isValid.IsValid(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.input, test.output)
		}
	}
}