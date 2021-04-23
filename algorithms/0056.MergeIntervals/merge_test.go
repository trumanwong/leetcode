package main

import (
	"leetcode/algorithms/0056.MergeIntervals/merge"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		output    [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}

	for _, test := range tests {
		ret := merge.Merge(test.intervals)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input %v; expected %v", ret, test.intervals, test.output)
				break
			}
		}
	}
}
