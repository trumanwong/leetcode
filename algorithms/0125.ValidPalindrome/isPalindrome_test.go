package main

import (
	"leetcode/algorithms/0125.ValidPalindrome/isPalindrome"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}
	for _, test := range tests {
		ret := isPalindrome.IsPalindrome(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.input, test.output)
		}
	}
}
