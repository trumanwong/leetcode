package main

import (
	"testing"
	"truman.com/leetcode/55.JumpGame/canJump"
)

func TestCanJump(t *testing.T)  {
	tests := []struct{
		input []int
		output bool
	}{
		{[]int{2,3,1,1,4}, true},
		{[]int{3,2,1,0,4}, false},
	}
	for _, test := range tests {
		ret := canJump.CanJump(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}