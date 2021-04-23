package main

import (
	"leetcode/algorithms/1156.SwapForLongestRepeatedCharacterSubstring/maxRepOpt1"
	"testing"
)

func TestMaxRepOpt1(t *testing.T) {
	tests := []struct {
		date   string
		output int
	}{
		{"ababa", 3},
		{"aaabaaa", 6},
		{"aaabbaaa", 4},
		{"aaaaa", 5},
		{"abcdef", 1},
	}

	for _, test := range tests {
		ret := maxRepOpt1.MaxRepOpt1(test.date)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.date, test.output)
		}
	}
}
