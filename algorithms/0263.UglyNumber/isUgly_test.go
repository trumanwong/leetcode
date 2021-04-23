package main

import (
	"leetcode/algorithms/0263.UglyNumber/isUgly"
	"testing"
)

func TestIsUgly(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{1, true},
		{6, true},
		{8, true},
		{14, false},
	}
	for _, test := range tests {
		ret := isUgly.IsUgly(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}
