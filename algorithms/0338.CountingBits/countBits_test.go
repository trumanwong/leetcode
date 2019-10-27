package main

import (
	"leetcode/algorithms/0338.CountingBits/countBits"
	"testing"
)

func TestCountBits(t *testing.T)  {
	tests := []struct{
		input int
		output []int
	}{
		{2,[]int{0,1,1}},
		{5,[]int{0,1,1,2,1,2}},
	}

	for _, test := range tests {
		ret := countBits.CountBits(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %d; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}