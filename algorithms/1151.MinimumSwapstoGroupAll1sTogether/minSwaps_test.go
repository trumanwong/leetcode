package main

import (
	"leetcode/algorithms/1151.MinimumSwapstoGroupAll1sTogether/minSwaps"
	"testing"
)

func TestMinSwaps(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{1,0,1,0,1}, 1},
		{[]int{0,0,0,1,0}, 0},
		{[]int{1,0,1,0,1,0,0,1,1,0,1}, 3},
	}

	for _, test := range tests {
		ret := minSwaps.MinSwaps(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}