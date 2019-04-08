package main

import (
	"testing"
	"truman.com/leetcode/202.HappyNumber/isHappy"
)

func TestIsHappy(t *testing.T)  {
	tests := []struct{
		input int
		output bool
	}{
		{19,true},
	}

	for _, test := range tests {
		ret := isHappy.IsHappy(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}