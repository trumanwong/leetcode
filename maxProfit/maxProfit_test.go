package main

import (
	"testing"
	"truman.com/leetcode/maxProfit/maxProfit"
)

func TestMaxProfit(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	} {
		{[]int{7,1,5,3,6,4}, 5},
		{[]int{7,6,4,3,1}, 0},
	}

	for _, test := range tests {
		ret := maxProfit.MaxProfit(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}