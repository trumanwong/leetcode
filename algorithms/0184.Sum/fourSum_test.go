package main

import (
	"leetcode/algorithms/0184.Sum/fourSum"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}

	for _, test := range tests {
		ret := fourSum.FourSum(test.nums, test.target)
		if len(ret) != len(test.output) {
			t.Errorf("Got %v for input %v, target = %d; expected %v", ret, test.nums, test.target, test.output)
		}
		for i, arr := range ret {
			judge := true
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}

			if !judge {
				t.Errorf("Got %v for input %v, target = %d; expected %v", ret, test.nums, test.target, test.output)
			}
		}
	}
}
