package main

import (
	"leetcode/algorithms/1137.N-thTribonacciNumber/tribonacci"
	"testing"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		N      int
		output int
	}{
		{4, 4},
		{25, 1389537},
	}

	for _, test := range tests {
		ret := tribonacci.Tribonacci(test.N)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.N, test.output)
		}

	}
}
