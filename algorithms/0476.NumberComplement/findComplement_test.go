package main

import (
	"leetcode/algorithms/0476.NumberComplement/findComplement"
	"testing"
)

func TestFindComplement(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{5,2},
		{1,0},
	}

	for _, test := range tests {
		ret := findComplement.FindComplement(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}