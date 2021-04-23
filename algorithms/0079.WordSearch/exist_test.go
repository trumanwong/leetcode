package main

import (
	"leetcode/algorithms/0079.WordSearch/exist"
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		output bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}

	for _, test := range tests {
		ret := exist.Exist(test.board, test.word)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.board, test.output)
		}
	}
}
