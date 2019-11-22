package main

import (
	"leetcode/algorithms/0999.AvailableCapturesforRook/numRookCaptures"
	"testing"
)

func TestNumRookCaptures(t *testing.T)  {
	tests := []struct{
		board [][]byte
		output int
	}{
		{[][]byte{{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.','.'},{'.','.','.','R','.','.','.','p'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'}}, 3},
	}

	for _, test := range tests {
		ret := numRookCaptures.NumRookCaptures(test.board)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.board, test.output)
		}
	}
}