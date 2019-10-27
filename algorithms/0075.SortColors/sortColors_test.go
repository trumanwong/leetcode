package main

import (
	"leetcode/algorithms/0075.SortColors/sortColors"
	"testing"
)

func TestSortColors(t *testing.T)  {
	tests := []struct {
		input []int
		output []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		temp := test.input
		sortColors.SortColors(test.input)
		length := len(test.input)
		for i := 0; i < length; i++ {
			if test.input[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", test.input, temp, test.output)
				break
			}
		}
	}
}
