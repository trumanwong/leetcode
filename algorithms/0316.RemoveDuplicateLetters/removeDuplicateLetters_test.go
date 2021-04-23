package main

import (
	"leetcode/algorithms/0316.RemoveDuplicateLetters/removeDuplicateLetters"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
	}

	for _, test := range tests {
		ret := removeDuplicateLetters.RemoveDuplicateLetters(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
