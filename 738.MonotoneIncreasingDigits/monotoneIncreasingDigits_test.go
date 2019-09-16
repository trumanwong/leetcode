package main

import (
	"testing"
	"truman.com/leetcode/738.MonotoneIncreasingDigits/monotoneIncreasingDigits"
)

func TestMonotoneIncreasingDigits(t *testing.T)  {
	tests := []struct{
		N int
		output int
	}{
		{10,9},
		{1234,1234},
		{332,299},
	}

	for _, test := range tests {
		ret := monotoneIncreasingDigits.MonotoneIncreasingDigits(test.N)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.N, test.output)
		}
	}
}