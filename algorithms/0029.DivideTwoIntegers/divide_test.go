package main

import (
	"leetcode/algorithms/0029.DivideTwoIntegers/divide"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		result   int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{1, -1, -1},
		{-2147483648, -1, 2147483647},
	}

	for _, test := range tests {
		ret := divide.Divide(test.dividend, test.divisor)
		if ret != test.result {
			t.Errorf("Got %d for input %d / %d; expected %d", ret, test.dividend, test.divisor, test.result)
		}
	}
}
