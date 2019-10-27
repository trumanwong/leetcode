package main

import (
	"leetcode/algorithms/0357.CountNumberswithUniqueDigits/countNumbersWithUniqueDigits"
	"testing"
)

func TestCountNumbersWithUniqueDigits(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{2,91},
	}

	for _, test := range tests {
		ret := countNumbersWithUniqueDigits.CountNumbersWithUniqueDigits(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}