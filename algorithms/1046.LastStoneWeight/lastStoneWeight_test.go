package main

import (
	"leetcode/algorithms/1046.LastStoneWeight/lastStoneWeight"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{3, 7, 2}, 2},
		{[]int{7, 6, 7, 6, 9}, 3},
	}

	for _, test := range tests {
		ret := lastStoneWeight.LastStoneWeight(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
