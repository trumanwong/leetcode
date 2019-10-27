package main

import (
	"leetcode/algorithms/0201.BitwiseANDofNumbersRange/rangeBitwiseAnd"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T)  {
	tests := []struct{
		m int
		n int
		output int
	}{
		{5,7,4},
		{0,1,0},
	}

	for _, test := range tests {
		ret := rangeBitwiseAnd.RangeBitwiseAnd(test.m, test.n)
		if ret != test.output {
			t.Errorf("Got %d for input (%d, %d); expected %d", ret, test.m, test.n, test.output)
		}

	}
}