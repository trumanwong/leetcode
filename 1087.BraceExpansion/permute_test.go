package main

import (
	"testing"
	"truman.com/leetcode/1087.BraceExpansion/permute"
)

func TestPermute(t *testing.T)  {
	tests := []struct{
		input string
		output []string
	}{
		{"{a,b}c{d,e}f", []string{"acdf","acef","bcdf","bcef"}},
		{"abcd", []string{"abcd"}},
		{"{a,b}{z,x,y}", []string{"ax","ay","az","bx","by","bz"}},
	}

	for _, test := range tests {
		ret := permute.Permute(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s; expected %v", ret, test.input, test.output)
			}
		}
	}
}