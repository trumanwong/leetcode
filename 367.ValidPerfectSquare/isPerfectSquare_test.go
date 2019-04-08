package main

import (
	"testing"
	"truman.com/leetcode/367.ValidPerfectSquare/isPerfectSquare"
)

func TestIsPerfectSquare(t *testing.T)  {
	tests := []struct{
		input int
		output bool
	}{
		{16,true},
		{14,false},
	}
	for _, test := range tests {
		ret := isPerfectSquare.IsPerfectSquare(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}