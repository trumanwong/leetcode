package main

import (
	"testing"
	"truman.com/leetcode/73.SetMatrixZeroes/setZeroes"
)

func TestGetZeroes(t *testing.T)  {
	tests := []struct{
		input [][]int
		matrix [][]int
		output [][]int
	}{
		{[][]int{{1,1,1}, {1,0,1}, {1,1,1}},
			[][]int{{1,1,1}, {1,0,1}, {1,1,1}},
			[][]int{{1,0,1},{0,0,0},{1,0,1}},
		},
		{[][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
			[][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
			[][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
		},
	}

	for _, test := range tests {
		setZeroes.GetZeroes(test.input)
		for i, v1 := range test.input {
			for j, v2 := range v1 {
				if v2 != test.output[i][j] {
					t.Errorf("Got %v for input %v; expected %v", test.input, test.matrix, test.output)
					break
				}
			}
		}
	}
}