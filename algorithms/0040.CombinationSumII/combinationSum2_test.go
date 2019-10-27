package main

import (
	"leetcode/algorithms/0040.CombinationSumII/combinationSum2"
	"testing"
)

func TestCombinationSum2(t *testing.T)  {
	tests := []struct{
		candidates []int
		target int
		output [][]int
	}{
		{[]int{10,1,2,7,6,1,5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}

	for _, test := range tests {
		ret := combinationSum2.CombinationSum2(test.candidates, test.target)
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