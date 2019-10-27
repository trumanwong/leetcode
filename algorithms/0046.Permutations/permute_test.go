package main

import (
	"leetcode/algorithms/0046.Permutations/permute"
	"testing"
)

func TestPermute(t *testing.T)  {
	tests := []struct {
		nums []int
		output [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1}}},
	}

	for _, test := range tests {
		ret := permute.Permute(test.nums)
		for i, arr := range ret {
			judge := true
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}

			if !judge {
				t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
				break
			}
		}
	}
}