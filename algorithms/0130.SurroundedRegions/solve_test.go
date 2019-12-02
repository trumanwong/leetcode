package main

import (
	"leetcode/algorithms/0130.SurroundedRegions/solve"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T)  {
	board := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
	input := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
	output := [][]byte{{'X','X','X','X'},{'X','X','X','X'},{'X','X','X','X'},{'X','O','X','X'}}
	solve.Solve(board)
	if !reflect.DeepEqual(board, output) {
		t.Errorf("Got %v for input %v; expected %v", board, input, output)
	}
}