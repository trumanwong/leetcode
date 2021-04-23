package main

import (
	"leetcode/algorithms/0136.SingleNumber/singleNumber"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, test := range tests {
		ret := singleNumber.SingleNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
