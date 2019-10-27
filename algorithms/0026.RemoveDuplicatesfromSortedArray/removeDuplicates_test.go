package main

import (
	"leetcode/algorithms/0026.RemoveDuplicatesfromSortedArray/removeDuplicates"
	"testing"
)

func TestRemoveDuplicates(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{1,1,2}, 2},
		{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
	}
	for _, test := range tests {
		ret := removeDuplicates.RemoveDuplicates(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}

	}
}
