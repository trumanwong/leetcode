package main

import (
	"leetcode/algorithms/0766.ToeplitzMatrix/isToeplitzMatrix"
	"testing"
)

func TestIsToeplitzMatrix(t *testing.T)  {
	tests := []struct{
		matrix [][]int
		output bool
	}{
		{[][]int{{1,2,3,4},
		{5,1,2,3},
		{9,5,1,2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}

	for _, test := range tests {
		ret := isToeplitzMatrix.IsToeplitzMatrix(test.matrix)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.matrix, test.output)
		}
	}
}