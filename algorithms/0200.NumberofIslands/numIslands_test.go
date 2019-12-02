package main

import (
	"leetcode/algorithms/0200.NumberofIslands/numIslands"
	"testing"
)

func TestNumIslands(t *testing.T)  {
	grid := [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	input := [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	output := 1
	ret := numIslands.NumIslands(grid)
	if ret != output {
		t.Errorf("Got %d for input %v; expected %d", ret, input, output)
	}
}