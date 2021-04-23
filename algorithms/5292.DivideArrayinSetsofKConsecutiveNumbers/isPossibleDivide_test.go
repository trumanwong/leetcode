package main

import (
	"leetcode/algorithms/5292.DivideArrayinSetsofKConsecutiveNumbers/isPossibleDivide"
	"testing"
)

func TestIsPossibleDivide(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output bool
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true},
		{[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true},
		{[]int{3, 3, 2, 2, 1, 1}, 3, true},
		{[]int{1, 2, 3, 4}, 3, false},
		{[]int{1, 2, 3, 4, 5, 6, 4, 4, 4}, 3, false},
		{[]int{15, 16, 17, 18, 19, 16, 17, 18, 19, 20, 6, 7, 8, 9, 10, 3, 4, 5, 6, 20}, 5, false},
	}

	for _, test := range tests {
		ret := isPossibleDivide.IsPossibleDivide(test.nums, test.k)
		if ret != test.output {
			t.Errorf("failure")
		}
	}
}
