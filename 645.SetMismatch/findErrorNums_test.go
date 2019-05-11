package main

import (
	"testing"
	"truman.com/leetcode/645.SetMismatch/findErrorNums"
)

func TestFindErrorNums(t *testing.T)  {
	tests := []struct {
		input []int
		output []int
	}{
		{[]int{1,2,2,4},[]int{2,3}},
	}

	for _, test := range tests {
		ret := findErrorNums.FindErrorNums(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
			}
		}
	}
}