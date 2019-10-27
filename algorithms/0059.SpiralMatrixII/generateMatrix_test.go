package main

import (
	"leetcode/algorithms/0059.SpiralMatrixII/generateMatrix"
	"testing"
)

func TestGenerateMatrix(t *testing.T)  {
	tests := []struct{
		n int
		output [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, test := range tests {
		ret := generateMatrix.GenerateMatrix(test.n)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input %d; expected %v", ret, test.n, test.output)
				break
			}
		}
	}
}