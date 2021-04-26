package main

import (
	"leetcode/algorithms/5738.SumofDigitsinBaseK/sumBase"
	"testing"
)

func TestSumBase(t *testing.T)  {
	tests := []struct{
		n int
		k int
		output int
	}{
		{34, 6, 9},
		{10, 10, 1},
	}

	for _, test := range tests {
		ret := sumBase.SumBase(test.n, test.k)
		if ret != test.output {
			t.Errorf("Got %d for input (%d, %d); expected %d", ret, test.n, test.k, test.output)
		}
	}
}