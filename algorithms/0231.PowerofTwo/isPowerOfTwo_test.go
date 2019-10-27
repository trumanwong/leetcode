package main

import (
	"leetcode/algorithms/0231.PowerofTwo/isPowerOfTwo"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T)  {
	tests := []struct{
		input int
		output bool
	}{
		{8,true},
		{121,false},
	}
	for _, test := range tests {
		ret := isPowerOfTwo.IsPowerOfTwo(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}