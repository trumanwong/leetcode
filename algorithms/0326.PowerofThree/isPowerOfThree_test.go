package main

import (
	"leetcode/algorithms/0326.PowerofThree/isPowerOfThree"
	"testing"
)

func TestIsPowerOfThree(t *testing.T)  {
	tests := []struct{
		n int
		output bool
	}{
		{3,true},
		{5,false},
	}
	for _, test := range tests {
		ret := isPowerOfThree.IsPowerOfThree(test.n)
		if test.output != ret {
			t.Errorf("Got %t for input %d; expected %t", ret, test.n, test.output)
		}
	}
}