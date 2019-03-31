package main

import (
	"testing"
	"truman.com/leetcode/152.MaximumProductSubarray/maxProduct"
)

func TestMaxProduct(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{2,3,-2,4},6},
		{[]int{-2,0,-1},0},
	}

	for _, test := range tests {
		ret := maxProduct.MaxProduct(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}