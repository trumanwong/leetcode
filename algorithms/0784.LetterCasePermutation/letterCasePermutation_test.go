package main

import (
	"leetcode/algorithms/0784.LetterCasePermutation/letterCasePermutation"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		S      string
		output []string
	}{
		{"a1b2", []string{"A1B2", "A1b2", "a1B2", "a1b2"}},
		{"3z4", []string{"3Z4", "3z4"}},
		{"12345", []string{"12345"}},
	}

	for _, test := range tests {
		ret := letterCasePermutation.LetterCasePermutation(test.S)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s; expected %s", ret, test.S, test.output)
				break
			}
		}
	}
}
