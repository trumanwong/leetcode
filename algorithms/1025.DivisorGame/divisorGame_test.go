package main

import (
	"leetcode/algorithms/1025.DivisorGame/divisorGame"
	"testing"
)

func TestDivisorGame(t *testing.T)  {
	tests := []struct{
		N int
		output bool
	}{
		{2,true},
		{3,false},
	}

	for _, test := range tests {
		ret := divisorGame.DivisorGame(test.N)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.N, test.output)
		}
	}
}