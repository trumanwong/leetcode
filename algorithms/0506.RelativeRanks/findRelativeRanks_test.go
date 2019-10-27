package main

import (
	"leetcode/algorithms/0506.RelativeRanks/findRelativeRanks"
	"testing"
)

func TestFindRelativeRanks(t *testing.T)  {
	tests := []struct{
		input []int
		output []string
	}{
		{[]int{5, 4, 3, 2, 1},[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	}

	for _, test := range tests {
		ret := findRelativeRanks.FindRelativeRanks(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}