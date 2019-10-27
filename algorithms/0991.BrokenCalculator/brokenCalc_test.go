package main

import (
	"leetcode/algorithms/0991.BrokenCalculator/brokenCalc"
	"testing"
)

func TestBrokenCalc(t *testing.T)  {
	tests := []struct{
		X int
		Y int
		output int
	}{
		{2,3,2},
		{5,8,2},
		{3,10,3},
		{1024,1,1023},
	}

	for _, test := range tests {
		ret := brokenCalc.BrokenCalc(test.X, test.Y)
		if ret != test.output {
			t.Errorf("Got %d for input X = %d, Y = %d; expected %d", ret, test.X, test.Y, test.output)
		}
	}
}