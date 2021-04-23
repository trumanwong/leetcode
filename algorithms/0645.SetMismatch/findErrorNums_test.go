package main

import (
	"leetcode/algorithms/0645.SetMismatch/findErrorNums"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
	}

	for _, test := range tests {
		ret := findErrorNums.FindErrorNums(test.nums)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
				break
			}
		}
	}
}
