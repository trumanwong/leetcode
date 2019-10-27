package main

import (
	"leetcode/algorithms/1017.ConverttoBase-2/baseNeg2"
	"testing"
)

func TestBaseNeg2(t *testing.T)  {
	tests := []struct{
		input int
		output string
	}{
		{2,"110"},
		{3,"111"},
		{4,"100"},
	}

	for _, test := range tests {
		ret := baseNeg2.BaseNeg2(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.input, test.output)
		}
	}
}