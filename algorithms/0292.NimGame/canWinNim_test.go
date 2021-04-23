package main

import (
	"leetcode/algorithms/0292.NimGame/canWinNim"
	"testing"
)

func TestCanWinNim(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{4, false},
		{5, true},
	}
	for _, test := range tests {
		ret := canWinNim.CanWinNim(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}
