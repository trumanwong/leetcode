package main

import (
	"testing"
	"truman.com/leetcode/14.LongestCommonPrefix/longestCommonPrefix"
)

func TestLongestCommonPrefix(t *testing.T)  {
	tests := []struct {
		input []string
		output string
	}{
		{[]string{"flower","flow","flight"}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
	}
	for _, test := range tests {
		ret := longestCommonPrefix.LongestCommonPrefix(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %v; expected %s", ret, test.input, test.output)
		}
	}
}
