package main

import (
	"testing"
	"truman.com/leetcode/17.LetterCombinationsofaPhoneNumber/letterCombinations"
)

func TestLetterCombinations(t *testing.T)  {
	tests := []struct{
		digits string
		output []string
	}{
		{"23", []string{"bd", "be", "bf", "ce", "cf", "ad", "ae", "af", "cd"}},
	}

	for _, test := range tests {
		ret := letterCombinations.LetterCombinations(test.digits)
		for i, v := range ret {
			if test.output[i] != v {
				t.Errorf("Got %v for input %s; expected %v", ret, test.digits, test.output)
				break
			}
		}
	}
}