package main

import (
	"leetcode/algorithms/1200.MinimumAbsoluteDifference/minimumAbsDifference"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T)  {
	tests := []struct{
		arr []int
		output [][]int
	}{
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{3,8,-10,23,19,-4,-14,27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
	}

	for _, test := range tests {
		ret := minimumAbsDifference.MinimumAbsDifference(test.arr)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input %v; expected %v", ret, test.arr, test.output)
				break
			}
		}
	}
}