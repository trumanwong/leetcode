package main

import (
	"leetcode/algorithms/0643.MaximumAverageSubarrayI/findMaxAverage"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	}

	for _, test := range tests {
		ret := findMaxAverage.FindMaxAverage(test.nums, test.k)
		if ret-test.output > 0.00001 {
			t.Errorf("Got %f for input nums = %v, k = %d; expected %f", ret, test.nums, test.k, test.output)
		}
	}
}
