package main

import (
	"leetcode/algorithms/0069.Sqrt/mySqrt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{4, 2},
		{8, 2},
	}
	for _, test := range tests {
		ret := mySqrt.MySqrt(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
