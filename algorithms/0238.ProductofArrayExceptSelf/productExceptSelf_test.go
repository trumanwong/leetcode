package main

import (
	"leetcode/algorithms/0238.ProductofArrayExceptSelf/productExceptSelf"
	"testing"
)

func TestProductExceptSelf(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{1,2,3,4},[]int{24,12,8,6}},
	}

	for _, test := range tests {
		ret := productExceptSelf.ProductExceptSelf(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
			}
		}
	}
}