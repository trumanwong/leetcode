package main

import (
	"leetcode/algorithms/1217.PlaywithChips/minCostToMoveChips"
	"testing"
)

func TestMinCostToMoveChips(t *testing.T)  {
	tests := []struct{
		chips []int
		output int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{2, 2, 2, 3, 3}, 2},
	}

	for _, test := range tests {
		ret := minCostToMoveChips.MinCostToMoveChips(test.chips)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.chips, test.output)
		}
	}
}