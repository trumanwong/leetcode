package main

import (
	"leetcode/algorithms/0977.SquaresofaSortedArray/sortedSquares"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, test := range tests {
		ret := sortedSquares.SortedSquares(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; exptected %v", ret, test.input, test.output)
				break
			}
		}
	}
}
