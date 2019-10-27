package main

import (
	"leetcode/algorithms/1079.LetterTilePossibilities/numTilePossibilities"
	"testing"
)

func TestNumTilePossibilities(t *testing.T)  {
	tests := []struct{
		input string
		output int
	}{
		{"AAB", 8},
		{"AAABBC", 188},
	}

	for _, test := range tests {
		ret := numTilePossibilities.NumTilePossibilities(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}