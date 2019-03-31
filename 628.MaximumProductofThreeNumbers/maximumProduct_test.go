package main

import (
	"testing"
	"truman.com/leetcode/628.MaximumProductofThreeNumbers/maximumProduct"
)

func TestMaximumProduct(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{1,2,3},6},
		{[]int{1,2,3,4},24},
	}
	for _, test := range tests {
		ret := maximumProduct.MaximumProduct(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}