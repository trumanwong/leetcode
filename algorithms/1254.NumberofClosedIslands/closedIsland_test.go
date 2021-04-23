package main

import (
	"leetcode/algorithms/1254.NumberofClosedIslands/closedIsland"
	"testing"
)

func TestClosedIsland(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}, 2},
	}

	for _, test := range tests {
		ret := closedIsland.ClosedIsland(test.grid)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.grid, test.output)
		}
	}
}
