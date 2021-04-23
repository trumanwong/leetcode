package main

import (
	"leetcode/algorithms/0633.SumofSquareNumbers/judgeSquareSum"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{5, true},
		{3, false},
	}
	for _, test := range tests {
		ret := judgeSquareSum.JudgeSquareSum(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}
