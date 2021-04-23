package main

import (
	"leetcode/algorithms/0713.SubarrayProductLessThanK/numSubarrayProductLessThanK"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
	}

	for _, test := range tests {
		ret := numSubarrayProductLessThanK.NumSubarrayProductLessThanK(test.nums, test.k)
		if ret != test.output {
			t.Errorf("Got %d for nums = %v, k = %d; expected %d", ret, test.nums, test.k, test.output)
		}
	}
}
