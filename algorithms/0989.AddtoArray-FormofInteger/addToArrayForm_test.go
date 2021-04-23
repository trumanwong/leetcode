package main

import (
	"leetcode/algorithms/0989.AddtoArray-FormofInteger/addToArrayForm"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	tests := []struct {
		A      []int
		K      int
		output []int
	}{
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		ret := addToArrayForm.AddToArrayForm(test.A, test.K)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input A = %v, K = %d; expected %v", ret, test.A, test.K, test.output)
			}
		}
	}
}
