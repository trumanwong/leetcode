package main

import (
	"leetcode/algorithms/0229.MajorityElementII/majorityElement"
	"testing"
)

func TestMajorityElement(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{3,2,3},[]int{3}},
		{[]int{1,1,1,3,3,2,2,2},[]int{1,2}},
	}

	for _, test := range tests {
		ret := majorityElement.MajorityElement(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}