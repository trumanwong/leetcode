package main

import (
	"testing"
	"truman.com/leetcode/292.NimGame/canWinNim"
)

func TestCanWinNim(t *testing.T)  {
	tests := []struct{
		input int
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