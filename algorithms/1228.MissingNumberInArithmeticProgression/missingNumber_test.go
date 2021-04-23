package main

import (
	"leetcode/algorithms/1228.MissingNumberInArithmeticProgression/missingNumber"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		arr    []int
		output int
	}{
		{[]int{5, 7, 11, 13}, 9},
		{[]int{15, 13, 12}, 14},
	}

	for _, test := range tests {
		ret := missingNumber.MissingNumber(test.arr)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.arr, test.output)
		}
	}
}
