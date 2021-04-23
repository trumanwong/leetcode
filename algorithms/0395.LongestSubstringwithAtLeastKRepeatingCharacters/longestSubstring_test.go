package main

import (
	"leetcode/algorithms/0395.LongestSubstringwithAtLeastKRepeatingCharacters/longestSubstring"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output int
	}{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
	}

	for _, test := range tests {
		ret := longestSubstring.LongestSubstring(test.s, test.k)
		if ret != test.output {
			t.Errorf("Got %d for input s = %s, k = %d; expected %d", ret, test.s, test.k, test.output)
		}
	}
}
