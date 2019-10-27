package main

import (
	"leetcode/algorithms/0066.PlusOne/plusOne"
	"testing"
)

func TestPlushOne(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	} {
		{[]int{1,2,3}, []int{1,2,4}},
		{[]int{4,3,2,1}, []int{4,3,2,2}},
		{[]int{9,9,9,9}, []int{1,0,0,0,0}},
	}

	for _, test := range tests {
		ret := plusOne.PlusOne(test.input)
		length := len(ret)
		for i := 0; i < length; i++ {
			if ret[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}