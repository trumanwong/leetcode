package main

import (
	"leetcode/algorithms/0922.SortArrayByParityII/sortArrayByParityII"
	"testing"
)

func TestSortArrayByParityII(t *testing.T) {
	tests := []struct {
		A      []int
		output []int
	}{
		{[]int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
	}
	for _, test := range tests {
		A := make([]int, len(test.A))
		copy(A, test.A)
		ret := sortArrayByParityII.SortArrayByParityII(A)
		for i := 0; i < len(A); i++ {
			if A[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.A, test.output)
			}
		}
	}
}
