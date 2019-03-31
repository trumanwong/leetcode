package main

import (
	"testing"
	"truman.com/leetcode/1012.ComplementofBase10Integer/bitwiseComplement"
)

func TestBitwiseComplement(t *testing.T)  {
	tests := []struct{
		N int
		output int
	}{
		{5,2},
		{7,0},
		{10,5},
	}
	for _, test := range tests {
		ret := bitwiseComplement.BitwiseComplement(test.N)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.N, test.output)
		}
	}
}