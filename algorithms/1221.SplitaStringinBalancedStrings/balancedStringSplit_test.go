package main

import (
	"leetcode/algorithms/1221.SplitaStringinBalancedStrings/balancedStringSplit"
	"testing"
)

func TestBalancedStringSplit(t *testing.T)  {
	tests := []struct{
		s string
		output int
	}{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
	}

	for _, test := range tests {
		ret := balancedStringSplit.BalancedStringSplit(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}