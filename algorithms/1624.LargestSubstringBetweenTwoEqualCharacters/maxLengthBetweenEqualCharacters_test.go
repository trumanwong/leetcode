package main

import (
	"leetcode/algorithms/1624.LargestSubstringBetweenTwoEqualCharacters/maxLengthBetweenEqualCharacters"
	"testing"
)

func TestMaxLengthBetweenEqualCharacters(t *testing.T)  {
	tests := []struct{
		s string
		output int
	}{
		{"aa", 0},
		{"abca", 2},
		{"cbzxy", -1},
		{"cabbac", 4},
	}
	for _, test := range tests {
		ret := maxLengthBetweenEqualCharacters.MaxLengthBetweenEqualCharacters(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}