package main

import (
	"leetcode/algorithms/0434.NumberofSegmentsinaString/countSegments"
	"testing"
)

func TestCountSegments(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"Hello, my name is John", 5},
	}

	for _, test := range tests {
		ret := countSegments.CountSegments(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}
