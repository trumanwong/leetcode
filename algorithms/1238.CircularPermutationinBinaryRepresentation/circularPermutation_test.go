package main

import (
	"leetcode/algorithms/1238.CircularPermutationinBinaryRepresentation/circularPermutation"
	"testing"
)

func TestCircularPermutation(t *testing.T) {
	tests := []struct {
		n      int
		start  int
		output []int
	}{
		{2, 3, []int{3, 2, 0, 1}},
		{3, 2, []int{2, 3, 1, 0, 4, 5, 7, 6}},
	}

	for _, test := range tests {
		ret := circularPermutation.CircularPermutation(test.n, test.start)
		judge := true
		for i, v := range ret {
			if v != test.output[i] {
				judge = false
				break
			}
		}
		if !judge {
			t.Errorf("Got %v for input n = %d, start = %d; expected %v", ret, test.n, test.start, test.output)
			break
		}
	}
}
