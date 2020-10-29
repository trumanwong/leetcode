package main

import (
	"leetcode/algorithms/1550.ThreeConsecutiveOdds/threeConsecutiveOdds"
	"testing"
)

func TestThreeConsecutiveOdds(t *testing.T)  {
	tests := []struct{
		arr []int
		output bool
	} {
		{[]int{2, 6, 4, 1}, false},
		{[]int{1,2,34,3,4,5,7,23,12}, true},
	}

	for _, test := range tests {
		ret := threeConsecutiveOdds.ThreeConsecutiveOdds(test.arr)
		if test.output != ret {
			t.Errorf("Got %t for input %v; expected %t", ret, test.arr, test.output)
		}
	}
}