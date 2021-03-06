package main

import (
	"leetcode/algorithms/0005.LongestPalindromicSubstring/longestPalindrome"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"bb", "bb"},
	}

	for _, test := range tests {
		ret := longestPalindrome.LongestPalindrome(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}
