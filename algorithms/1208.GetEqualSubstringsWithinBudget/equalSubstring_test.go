package main

import (
	"leetcode/algorithms/1208.GetEqualSubstringsWithinBudget/equalSubstring"
	"testing"
)

func TestEqualSubstring(t *testing.T) {
	tests := []struct {
		s       string
		t       string
		maxCost int
		output  int
	}{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
	}

	for _, test := range tests {
		ret := equalSubstring.EqualSubstring(test.s, test.t, test.maxCost)
		if ret != test.output {
			t.Errorf("Got %d for s = %s, t = %s, maxCost = %d; expected %d", ret, test.s, test.t, test.maxCost, test.output)
		}
	}
}
