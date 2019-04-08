package main

import (
	"testing"
	"truman.com/leetcode/191.Numberof1Bits/hammingWeight"
)

func TestHammingWeight(t *testing.T)  {
	tests := []struct{
		input uint32
		output int
	}{
		{11,3},
		{128,1},
	}
	for _, test := range tests {
		ret := hammingWeight.HammingWeight(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}