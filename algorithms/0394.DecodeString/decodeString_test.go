package main

import (
	"leetcode/algorithms/0394.DecodeString/decodeString"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, test := range tests {
		ret := decodeString.DecodeString(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
