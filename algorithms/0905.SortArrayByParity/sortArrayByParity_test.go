package main

import (
	"leetcode/algorithms/0905.SortArrayByParity/sortArrayByParity"
	"testing"
)

func TestSortArrayByParity(t *testing.T)  {
	tests := []struct{
		A []int
		output []int
	}{
		{[]int{3,1,2,4}, []int{2,4,3,1}},
		{[]int{1,3}, []int{1,3}},
		{[]int{0,1,2}, []int{0, 2, 1}},
		{[]int{1, 3, 5}, []int{1,3,5}},
		{[]int{1,0,3}, []int{0,1,3}},
		{[]int{0,2,1,4}, []int{0,2,4,1}},
		{[]int{3,0,5,1}, []int{0,3,5,1}},
		{[]int{0,2,4,1,6,8},[]int{0,2,4,6,8,1}},
	}

	for _, test := range tests {
		A := make([]int, len(test.A))
		copy(A, test.A)
		ret := sortArrayByParity.SortArrayByParity(A)
		for i := 0; i < len(ret); i++ {
			if ret[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.A, test.output)
				break
			}
		}
	}
}