package main

import (
	"testing"
	"truman.com/leetcode/1014.BestSightseeingPair/maxScoreSightseeingPair"
)

func TestMaxScoreSightseeingPair(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{8,1,5,2,6},11},
	}
	for _, test := range tests {
		ret := maxScoreSightseeingPair.MaxScoreSightseeingPair(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}