package main

import (
	"testing"
	"truman.com/leetcode/459.RepeatedSubstringPattern/repeatedSubstringPattern"
)

func TestRepeatedSubstringPattern(t *testing.T)  {
	tests := []struct{
		input string
		output bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	}

	for _, test := range tests {
		ret := repeatedSubstringPattern.RepeatedSubstringPattern(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}