package main

import (
	"leetcode/algorithms/0279.PerfectSquares/numSquares"
	"testing"
)

func TestNumSquares(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{12, 3},
		{13,2},
	}

	for _, test := range tests {
		ret := numSquares.NumSquares(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}