package main

import (
	"leetcode/algorithms/0137.SingleNumberII/singleNumber"
	"testing"
)

func TestSingleNumber(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{2,2,3,2}, 3},
		{[]int{0,1,0,1,0,1,99}, 99},
	}

	for _, test := range tests {
		ret := singleNumber.SingleNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}