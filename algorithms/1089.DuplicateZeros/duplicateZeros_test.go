package main

import (
	"leetcode/algorithms/1089.DuplicateZeros/duplicateZeros"
	"testing"
)

func TestDuplicateZeros(t *testing.T)  {
	tests := []struct{
		input []int
		origin []int
		output []int
	}{
		{[]int{1,0,2,3,0,4,5,0}, []int{1,0,2,3,0,4,5,0},[]int{1,0,0,2,3,0,0,4}},
		{[]int{1,2,3}, []int{1,2,3},[]int{1,2,3}},
		{[]int{8,4,5,0,0,0,0,7}, []int{8,4,5,0,0,0,0,7},[]int{8,4,5,0,0,0,0,0}},
	}

	for _, test := range tests {
		duplicateZeros.DuplicateZeros(test.input)
		for i, v := range test.input {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", test.input, test.origin, test.output)
				break
			}
		}
	}
}