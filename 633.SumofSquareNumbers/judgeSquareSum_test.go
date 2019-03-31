package main

import (
	"testing"
	"truman.com/leetcode/633.SumofSquareNumbers/judgeSquareSum"
)

func TestJudgeSquareSum(t *testing.T)  {
	tests := []struct{
		input int
		output bool
	}{
		{5,true},
		{3,false},
	}
	for _, test := range tests {
		ret := judgeSquareSum.JudgeSquareSum(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}