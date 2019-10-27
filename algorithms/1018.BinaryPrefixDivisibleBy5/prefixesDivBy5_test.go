package main

import (
	"leetcode/algorithms/1018.BinaryPrefixDivisibleBy5/prefixesDivBy5"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T)  {
	tests := []struct{
		input []int
		output []bool
	}{
		{[]int{0,1,1},[]bool{true,false,false}},
		{[]int{1,1,1},[]bool{false,false,false}},
		{[]int{0,1,1,1,1,1},[]bool{true,false,false,false,true,false}},
		{[]int{1,1,1,0,1},[]bool{false,false,false,false,false}},
		{[]int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0},
			[]bool{false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true}},
	}
	for _, test := range tests {
		ret := prefixesDivBy5.PrefixesDivBy5(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
			}
		}
	}
}