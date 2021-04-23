package main

import (
	"leetcode/algorithms/0696.CountBinarySubstrings/countBinarySubstrings"
	"testing"
)

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"00110011", 6},
		{"10101", 4},
	}

	for _, test := range tests {
		ret := countBinarySubstrings.CountBinarySubstrings(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}
