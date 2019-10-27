package main

import (
	"leetcode/algorithms/0007.ReverseInteger/reverse"
	"testing"
)

func TestReverse(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120,21},
	}
	for _, test := range tests {
		ret := reverse.Reverse(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}