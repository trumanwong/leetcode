package main

import (
	"leetcode/algorithms/1099.TwoSumLessThanK/twoSumLessThanK"
	"testing"
)

func TestTwoSumLessThanK(t *testing.T)  {
	tests := []struct{
		A []int
		K int
		Output int
	}{
		{[]int{34,23,1,24,75,33,54,8}, 60, 58},
		{[]int{10,20,30}, 15,-1},
	}

	for _, test := range tests {
		ret := twoSumLessThanK.TwoSumLessThanK(test.A, test.K)
		if ret != test.Output {
			t.Errorf("Got %d for input A = %v, K = %d; expected %d", ret, test.A, test.K, test.Output)
		}
	}
}