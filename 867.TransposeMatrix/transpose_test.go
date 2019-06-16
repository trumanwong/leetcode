package main

import (
	"testing"
	"truman.com/leetcode/867.TransposeMatrix/transpose"
)

func TestTranspose(t *testing.T)  {
	tests := []struct{
		input [][]int
		output [][]int
	}{
		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{1,4,7},{2,5,8},{3,6,9}}},
		{[][]int{{1,2,3},{4,5,6}}, [][]int{{1,4},{2,5},{3,6}}},
	}

	for _, test := range tests {
		ret := transpose.Transpose(test.input)
		for i := 0; i < len(ret); i++ {
			for j := 0; j < len(ret[i]); j++ {
				if ret[i][j] != test.output[i][j] {
					t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
					break
				}
			}
		}
	}
}