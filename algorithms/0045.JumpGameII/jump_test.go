package main

import (
	"leetcode/algorithms/0045.JumpGameII/jump"
	"testing"
)

func TestJump(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{2,3,1,1,4}, 2},
	}
	for _, test := range tests {
		ret := jump.Jump(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}