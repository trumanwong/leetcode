package main

import (
	"leetcode/algorithms/1119.RemoveVowelsfromaString/removeVowels"
	"testing"
)

func TestRemoveVowels(t *testing.T) {
	tests := []struct {
		S      string
		output string
	}{
		{"aaabbbiiiwww", "bbbwww"},
	}

	for _, test := range tests {
		ret := removeVowels.RemoveVowels(test.S)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.S, test.output)
		}
	}
}
