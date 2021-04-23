package main

import (
	"leetcode/algorithms/0521.LongestUncommonSubsequenceI/findLUSlength"
	"testing"
)

func TestFindLUSlength(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output int
	}{
		{"aba", "cdc", 3},
	}

	for _, test := range tests {
		ret := findLUSlength.FindLUSlength(test.a, test.b)
		if ret != test.output {
			t.Errorf("Got %d for input a = %s, b = %s; expected %d", ret, test.a, test.b, test.output)
		}
	}
}
