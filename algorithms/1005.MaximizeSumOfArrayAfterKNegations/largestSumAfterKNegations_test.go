package main

import (
	"leetcode/algorithms/1005.MaximizeSumOfArrayAfterKNegations/largestSumAfterKNegations"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	tests := []struct {
		A      []int
		K      int
		output int
	}{
		{[]int{4, 2, 3}, 1, 5},
		{[]int{3, -1, 0, 2}, 3, 6},
		{[]int{2, -3, -1, 5, -4}, 2, 13},
	}

	for _, test := range tests {
		ret := largestSumAfterKNegations.LargestSumAfterKNegations(test.A, test.K)
		if ret != test.output {
			t.Errorf("Got %d for input %v, k = %d; expected %d", ret, test.A, test.K, test.output)
		}
	}
}
