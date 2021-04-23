package main

import (
	"leetcode/algorithms/0010.RegularExpressionMatching/isMatch"
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s      string
		p      string
		output bool
	}{
		{"aa", "a", false},
		{"ab", ".*", true},
		{"aa", "a*", true},
		{"ab", ".*c", false},
		{"a", ".*..a*", false},
	}

	for _, test := range tests {
		ret := isMatch.IsMatch(test.s, test.p)
		if ret != test.output {
			t.Errorf("Got %t for s = %s, p = %s; expected %t", ret, test.s, test.p, test.output)
		}
	}
}
