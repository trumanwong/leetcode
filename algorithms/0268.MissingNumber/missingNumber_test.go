package main

import (
	"leetcode/algorithms/0268.MissingNumber/missingNumber"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for _, test := range tests {
		ret := missingNumber.MissingNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
