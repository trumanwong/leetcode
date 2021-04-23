package main

import (
	"leetcode/algorithms/0172.FactorialTrailingZeroes/trailingZeroes"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{3, 0},
		{5, 1},
	}
	for _, test := range tests {
		ret := trailingZeroes.TrailingZeroes(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}
