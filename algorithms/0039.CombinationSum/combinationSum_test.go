package main

import (
	"leetcode/algorithms/0039.CombinationSum/combinationSum"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	}

	for _, test := range tests {
		ret := combinationSum.CombinationSum(test.candidates, test.target)
		for i, arr := range ret {
			judge := true
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}

			if !judge {
				t.Errorf("Got %v for input %v, target = %d; expected %v", ret, test.candidates, test.target, test.output)
				break
			}
		}
	}
}
