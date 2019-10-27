package main

import (
	"leetcode/algorithms/0679.24Game/judgePoint24"
	"testing"
)

func TestJudgePoint24(t *testing.T)  {
	tests := []struct{
		input []int
		output bool
	}{
		{[]int{4,1,8,7},true},
		{[]int{1,2,1,2},false},
		{[]int{3,3,8,8},false},
	}
	for _, test := range tests {
		ret := judgePoint24.JudgePoint24(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}