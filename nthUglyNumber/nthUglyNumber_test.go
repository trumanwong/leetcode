package main

import (
	"testing"
	"truman.com/leetcode/nthUglyNumber/nthUglyNumber"
)

func TestNthUglyNumber(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{1, 1},
		{2, 2},
		{10, 12},
	}
	for _, test := range tests {
		ret := nthUglyNumber.NthUglyNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}