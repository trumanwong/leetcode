package main

import (
	"leetcode/algorithms/0845.LongestMountaininArray/longestMountain"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	tests := []struct {
		A      []int
		output int
	}{
		{[]int{2, 1, 4, 7, 3, 2, 5}, 5},
		{[]int{2, 2, 2}, 0},
	}

	for _, test := range tests {
		ret := longestMountain.LongestMountain(test.A)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.A, test.output)
		}
	}
}
