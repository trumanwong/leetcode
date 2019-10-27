package main

import (
	"leetcode/algorithms/0459.RepeatedSubstringPattern/repeatedSubstringPattern"
	"testing"
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