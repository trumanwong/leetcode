package main

import (
	"leetcode/algorithms/0421.MaximumXORofTwoNumbersinanArray/findMaximumXOR"
	"testing"
)

func TestFindMaximumXOR(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{3, 10, 5, 25, 2, 8}, 28},
	}

	for _, test := range tests {
		ret := findMaximumXOR.FindMaximumXOR(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
