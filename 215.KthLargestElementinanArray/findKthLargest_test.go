package main

import (
	"testing"
	"truman.com/leetcode/215.KthLargestElementinanArray/findKthLargest"
)

func TestFindKthLargest(t *testing.T)  {
	tests := []struct{
		input []int
		k int
		output int
	}{
		{[]int{3,2,1,5,6,4}, 2, 5},
		{[]int{3,2,3,1,2,4,5,5,6},4,4},
	}

	for _, test := range tests {
		ret := findKthLargest.FindKthLargest(test.input, test.k)
		if ret != test.output {
			t.Errorf("Got %d for input %v, k = %d; expected %d", ret, test.input, test.k, test.output)
		}
	}
}
