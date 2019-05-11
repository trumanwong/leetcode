package main

import (
	"testing"
	"truman.com/leetcode/1025.DivisorGame/divisorGame"
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