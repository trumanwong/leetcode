package main

import (
	"leetcode/algorithms/0762.PrimeNumberofSetBitsinBinaryRepresentation/countPrimeSetBits"
	"testing"
)

func TestCountPrimeSetBits(t *testing.T) {
	tests := []struct {
		L      int
		R      int
		output int
	}{
		{6, 10, 4},
		{10, 15, 5},
	}

	for _, test := range tests {
		ret := countPrimeSetBits.CountPrimeSetBits(test.L, test.R)
		if ret != test.output {
			t.Errorf("Got %d for L=%d, R=%d; expected %d", ret, test.L, test.R, test.output)
		}
	}
}
