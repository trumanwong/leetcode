package main

import (
	"leetcode/algorithms/0409.LongestPalindrome/longestPalindrome"
	"testing"
)

func TestLongestPalindrome(t *testing.T)  {
	tests := []struct{
		s string
		output int
	}{
		{"abccccdd", 7},
	}

	for _, test := range tests {
		ret := longestPalindrome.LongestPalindrome(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}