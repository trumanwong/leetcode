package main

import (
	"leetcode/algorithms/1122.RelativeSortArray/relativeSortArray"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		arr1   []int
		arr2   []int
		output []int
	}{
		{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}

	for _, test := range tests {
		ret := relativeSortArray.RelativeSortArray(test.arr1, test.arr2)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for arr1 = %v, arr = %v; expected %v", ret, test.arr1, test.arr2, test.output)
				break
			}
		}
	}
}
