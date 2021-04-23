package main

import (
	"leetcode/algorithms/0003.LongestSubstringWithoutRepeatingCharacters/lengthOfLongestSubstring"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{"床前明月光，疑似地上霜", 11},
	}

	for _, test := range tests {
		ret := lengthOfLongestSubstring.LengthOfLongestSubstring(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}
