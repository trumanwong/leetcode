package main

import (
	"leetcode/algorithms/0542.01Matrix/updateMatrix"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		Input  [][]int
		Output [][]int
	}{
		{Input: [][]int{{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}},
			Output: [][]int{{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
		},
		{Input: [][]int{{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1}},
			Output: [][]int{{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1}}},
	}

	for _, test := range tests {
		ret := updateMatrix.UpdateMatrix(test.Input)
		for i := 0; i < len(ret); i++ {
			for j := 0; j < len(ret[i]); j++ {
				if ret[i][j] != test.Output[i][j] {
					t.Errorf("Got %v for input %v; expected %v", ret, test.Input, test.Output)
					break
				}
			}
		}
	}
}
