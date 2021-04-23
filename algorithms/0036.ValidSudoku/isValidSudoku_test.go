package main

import (
	"leetcode/algorithms/0036.ValidSudoku/isValidSudoku"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		board  [][]byte
		output bool
	}{
		{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, true},
	}

	for _, test := range tests {
		ret := isValidSudoku.IsValidSudoku(test.board)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.board, test.output)
		}
	}
}
