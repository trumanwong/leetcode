package main

import (
	"testing"
	"truman.com/leetcode/640.SolvetheEquation/solveEquation"
)

func TestSolveEquation(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"x+5-3+x=6+x-2", "x=2"},
		{"x=x", "Infinite solutions"},
		{"2x=x", "x=0"},
		{"2x+3x-6x=x+2", "x=-1"},
		{"x=x+2", "No solution"},
	}

	for _, test := range tests {
		ret := solveEquation.SolveEquation(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}