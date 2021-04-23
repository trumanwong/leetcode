package main

import (
	"leetcode/algorithms/1209.RemoveAllAdjacentDuplicatesinStringII/removeDuplicates"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output string
	}{
		{"abcd", 2, "abcd"},
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
	}

	for _, test := range tests {
		ret := removeDuplicates.RemoveDuplicates(test.s, test.k)
		if ret != test.output {
			t.Errorf("Got %s for input s = %s, k = %d; expected %s", ret, test.s, test.k, test.output)
		}
	}
}
