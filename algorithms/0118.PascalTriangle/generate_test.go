package main

import (
	"leetcode/algorithms/0118.PascalTriangle/generate"
	"testing"
)

func TestGenerate(t *testing.T)  {
	tests := []struct{
		input int
		output [][]int
	}{
		{5,[][]int{
			{1},{1,1},{1,2,1},{1,3,3,1},{1,4,6,4,1},
		}},
	}

	for _, test := range tests {
		ret := generate.Generate(test.input)
		for i := 0; i < len(ret); i++ {
			for j := 0; j < len(ret[i]); j++ {
				if ret[i][j] != test.output[i][j] {
					t.Errorf("Got %v for input %d; expected %v", ret, test.input, test.output)
					break
				}
			}
		}
	}
}