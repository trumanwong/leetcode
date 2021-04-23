package main

import (
	"leetcode/algorithms/0788.RotatedDigits/rotatedDigits"
	"testing"
)

func TestRotatedDigits(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{10, 4},
	}

	for _, test := range tests {
		ret := rotatedDigits.RotatedDigits(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
