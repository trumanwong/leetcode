package main

import (
	"leetcode/algorithms/0204.CountPrimes/countPrimes"
	"testing"
)

func TestCountPrimes(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{10, 4},
		{10000, 1229},
		{999983, 78497},
	}
	for _, test := range tests {
		ret := countPrimes.CountPrimes(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
