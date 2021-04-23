package main

import (
	"leetcode/algorithms/5272.CountServersthatCommunicate/countServers"
	"testing"
)

func TestCountServers(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 0}, {0, 1}}, 0},
		{[][]int{{1, 0}, {1, 1}}, 3},
		{[][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 4},
		{[][]int{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}}, 3},
	}

	for _, test := range tests {
		ret := countServers.CountServers(test.grid)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.grid, test.output)
		}
	}
}
