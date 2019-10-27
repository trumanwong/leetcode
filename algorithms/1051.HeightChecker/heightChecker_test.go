package main

import (
	"leetcode/algorithms/1051.HeightChecker/heightChecker"
	"testing"
)

func TestHeightChecker(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{1,1,4,2,1,3},3},
	}

	for _, test := range tests {
		ret := heightChecker.HeightChecker(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}