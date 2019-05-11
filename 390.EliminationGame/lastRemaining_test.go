package main

import (
	"testing"
	"truman.com/leetcode/390.EliminationGame/lastRemaining"
)

func TestLastRemaining(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{9,6},
	}

	for _, test := range tests {
		ret := lastRemaining.LastRemaining(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}