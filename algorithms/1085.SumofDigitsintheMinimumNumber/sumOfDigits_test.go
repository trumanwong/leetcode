package main

import (
	"leetcode/algorithms/1085.SumofDigitsintheMinimumNumber/sumOfDigits"
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{34, 23, 1, 24, 75, 33, 54, 8}, 0},
		{[]int{99, 77, 33, 66, 55}, 1},
	}

	for _, test := range tests {
		ret := sumOfDigits.SumOfDigits(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
