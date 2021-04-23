package main

import (
	"fmt"
	"leetcode/algorithms/0368.LargestDivisibleSubset/largestDivisibleSubset"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 3}, []int{1, 2}},
	}

	for _, test := range tests {
		ret := largestDivisibleSubset.LargestDivisibleSubset(test.nums)
		fmt.Println(ret)
	}
}
