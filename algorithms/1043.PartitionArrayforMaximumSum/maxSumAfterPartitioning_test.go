package main

import (
	"leetcode/algorithms/1043.PartitionArrayforMaximumSum/maxSumAfterPartitioning"
	"testing"
)

func TestMaxSumAfterPartitioning(t *testing.T)  {
	tests := []struct{
		A []int
		K int
		output int
	}{
		{[]int{1,15,7,9,2,5,10}, 3, 84},
	}

	for _, test := range tests {
		ret := maxSumAfterPartitioning.MaxSumAfterPartitioning(test.A, test.K)
		if ret != test.output {
			t.Errorf("Got %d for input %v, k = %d; expected %d", ret, test.A, test.K, test.output)
		}
	}
}