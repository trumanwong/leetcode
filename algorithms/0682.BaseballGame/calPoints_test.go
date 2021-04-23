package main

import (
	"leetcode/algorithms/0682.BaseballGame/calPoints"
	"testing"
)

func TestCalPoints(t *testing.T) {
	tests := []struct {
		ops    []string
		output int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}

	for _, test := range tests {
		ret := calPoints.CalPoints(test.ops)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.ops, test.output)
		}
	}
}
