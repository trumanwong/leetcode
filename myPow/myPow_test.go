package main

import (
	"math"
	"testing"
	"truman.com/leetcode/myPow/myPow"
)

func TestMyPow(t *testing.T)  {
	tests := []struct{
		input1 float64
		input2 int
		ans float64
	} {
		{2.10000, 3, 9.261000},
		{2.00000, -2, 0.250000},
		{2.00000, 10, 1024.000000},
	}

	for _, tt := range tests {
		ret := myPow.MyPow(tt.input1, tt.input2)
		if math.Dim(ret, tt.ans) < 0.000000 {
			t.Errorf("Got MyPow(%f, %d) for input %f; expected %f", tt.input1, tt.input2, tt.ans, ret)
		}
	}
}
