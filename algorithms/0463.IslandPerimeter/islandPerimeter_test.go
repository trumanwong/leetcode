package main

import (
	"leetcode/algorithms/0463.IslandPerimeter/islandPerimeter"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, 16},
	}

	for _, test := range tests {
		ret := islandPerimeter.IslandPerimeter(test.grid)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.grid, test.output)
		}
	}
}
