package main

import (
	"leetcode/algorithms/0258.AddDigits/addDigits"
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{38, 2},
	}
	for _, test := range tests {
		ret := addDigits.AddDigits(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
