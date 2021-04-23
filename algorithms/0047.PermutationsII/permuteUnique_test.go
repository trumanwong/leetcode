package main

import (
	"leetcode/algorithms/0047.PermutationsII/permuteUnique"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		nums   []int
		output [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}

	for _, test := range tests {
		ret := permuteUnique.PermuteUnique(test.nums)
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
