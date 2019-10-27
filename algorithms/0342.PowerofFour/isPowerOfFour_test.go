package main

import (
	"leetcode/algorithms/0342.PowerofFour/isPowerOfFour"
	"testing"
)

func TestIsPowerOfFour(t *testing.T)  {
	tests := []struct{
		n int
		output bool
	}{
		{4,true},
		{5,false},
		{7, false},
	}
	for _, test := range tests {
		ret := isPowerOfFour.IsPowerOfFour(test.n)
		if test.output != ret {
			t.Errorf("Got %t for input %d; expected %t", ret, test.n, test.output)
		}
	}
}