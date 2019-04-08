package main

import (
	"testing"
	"truman.com/leetcode/868.BinaryGap/binaryGap"
)

func TestBinaryGap(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{22,2},
		{5,2},
		{6,1},
		{8,0},
	}

	for _, test := range tests {
		ret := binaryGap.BinaryGap(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}