package main

import (
	"leetcode/algorithms/0414.ThirdMaximumNumber/thirdMax"
	"testing"
)

func TestThridMax(t *testing.T) {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{3,2,1},1},
		{[]int{1,2},2},
		{[]int{2,2,3,1},1},
		{[]int{1,2,-2147483648}, -2147483648},
		{[]int{1,-2147483648,2},-2147483648},
		{[]int{5,2,4,1,3,6,0},4},
	}

	for _, test := range tests {
		res := thirdMax.ThirdMax(test.input)
		if res != test.output {
			t.Errorf("Got %d for input %v; expected %d", res, test.input, test.output)
		}
	}
}
