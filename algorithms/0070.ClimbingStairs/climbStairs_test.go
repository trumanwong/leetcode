package main

import (
	"leetcode/algorithms/0070.ClimbingStairs/climbStairs"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}
	for _, test := range tests {
		ret := climbStairs.ClimbStairs(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
