package main

import (
	"leetcode/algorithms/0205.IsomorphicStrings/isIsomorphic"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		output bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"ab", "aa", false},
	}
	for _, test := range tests {
		ret := isIsomorphic.IsIsomorphic(test.s, test.t)
		if ret != test.output {
			t.Errorf("Got %t for s = %s, t = %s; expected %t", ret, test.s, test.t, test.output)
		}
	}
}
