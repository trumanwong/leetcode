package main

import (
	"testing"
	"truman.com/leetcode/1037.ValidBoomerang/isBoomerang"
)

func TestIsBoomerang(t *testing.T)  {
	tests := []struct{
		input [][]int
		output bool
	} {
		{[][]int{{1,1},{2,3},{3,2}}, true},
		{[][]int{{1,1},{2,2},{3,3}}, false},
	}

	for _, test := range tests {
		ret := isBoomerang.IsBoomerang(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}