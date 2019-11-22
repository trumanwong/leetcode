package main

import (
	"leetcode/algorithms/0343.IntegerBreak/integerBreak"
	"testing"
)

func TestIntegerBreak(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{2, 1},
		{10, 36},
	}

	for _, test := range tests {
		ret := integerBreak.IntegerBreak(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}