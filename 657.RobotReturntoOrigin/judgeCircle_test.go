package main

import (
	"testing"
	"truman.com/leetcode/657.RobotReturntoOrigin/judgeCircle"
)

func TestJudgeCircle(t *testing.T)  {
	tests := []struct{
		input string
		output bool
	}{
		{"UD",true},
		{"LL",false},
	}

	for _, test := range tests {
		ret := judgeCircle.JudgeCircle(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.input, test.output)
		}
	}
}