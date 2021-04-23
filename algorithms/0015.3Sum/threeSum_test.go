package main

import (
	"leetcode/algorithms/0015.3Sum/threeSum"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums   []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, test := range tests {
		ret := threeSum.ThreeSum(test.nums)
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
