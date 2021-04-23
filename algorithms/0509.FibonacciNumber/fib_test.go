package main

import (
	"leetcode/algorithms/0509.FibonacciNumber/fib"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{2, 1},
		{4, 3},
		{3, 2},
	}

	for _, test := range tests {
		ret := fib.Fib(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
