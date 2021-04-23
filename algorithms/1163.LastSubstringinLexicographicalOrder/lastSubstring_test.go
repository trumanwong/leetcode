package main

import (
	"leetcode/algorithms/1163.LastSubstringinLexicographicalOrder/lastSubstring"
	"testing"
)

func TestLastSubstring(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"leetcode", "tcode"},
		{"abab", "bab"},
	}

	for _, test := range tests {
		ret := lastSubstring.LastSubstring(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
