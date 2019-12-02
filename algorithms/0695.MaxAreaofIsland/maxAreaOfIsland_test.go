package main

import (
	"leetcode/algorithms/0695.MaxAreaofIsland/maxAreaOfIsland"
	"testing"
)

func TestNumIslands(t *testing.T)  {
	grid := [][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}}
	input := [][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}}
	output := 4
	ret := maxAreaOfIsland.MaxAreaOfIsland(grid)
	if ret != output {
		t.Errorf("Got %d for input %v; expected %d", ret, input, output)
	}
}