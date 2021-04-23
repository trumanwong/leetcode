package main

import (
	"leetcode/algorithms/0031.NextPermutation/nextPermutation"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		nextPermutation.NextPermutation(test.nums)
		for i, v := range test.nums {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", test.nums, nums, test.output)
			}
		}
	}
}
