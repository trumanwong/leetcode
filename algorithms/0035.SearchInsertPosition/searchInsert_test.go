package main

import (
	"leetcode/algorithms/0035.SearchInsertPosition/searchInsert"
	"testing"
)

func TestSearchInsert(t *testing.T)  {
	tests := []struct{
		input []int
		target int
		output int
	}{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
	}

	for _, test := range tests {
		ret := searchInsert.SearchInsert(test.input, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input %v, target %d; expected %d", ret, test.input, test.target, test.output)
		}
	}
}
