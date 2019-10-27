package main

import (
	"leetcode/algorithms/1073.AddingTwoNegabinaryNumbers/addNegabinary"
	"testing"
)

func TestAddNegabinary(t *testing.T)  {
	tests := []struct{
		arr1 []int
		arr2 []int
		output []int
	}{
		{[]int{1,1,1,1,1}, []int{1, 0, 1}, []int{1, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		ret := addNegabinary.AddNegabinary(test.arr1, test.arr2)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for %v + %v; expected %v", ret, test.arr1, test.arr2, test.output)
			}
		}
	}
}