package main

import (
	"testing"
	"truman.com/leetcode/680.ValidPalindromeII/validPalindrome"
)

func TestValidPalindrome(t *testing.T)  {
	tests := []struct{
		input string
		output bool
	}{
		{"aba", true},
		{"abca", true},
		{"acbba",true},
		{"acda",true},
		{"acdza",false},
		{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
		{"eeccccbebaeeabebccceea", false},
		{"eccer", true},
	}
	for _, test := range tests {
		ret := validPalindrome.ValidPalindrome(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.input, test.output)
		}
	}
}