package main

import (
	"testing"
	"truman.com/leetcode/isPalindrome/isPalindrome"
)

func TestIsPalindrome(t *testing.T)  {
	tests := []struct{
		input int
		ans bool
	} {
		{10, false},
		{-121, false},
		{121, true},
	}

	for _, tt := range tests {
		ret := isPalindrome.IsPalindrome(tt.input)
		if ret != tt.ans {
			t.Errorf("Got %t for input %d; expected %t", ret, tt.input, tt.ans)
		}
	}
}
