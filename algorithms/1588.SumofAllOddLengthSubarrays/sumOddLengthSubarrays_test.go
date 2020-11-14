package main

import (
	"leetcode/algorithms/1588.SumofAllOddLengthSubarrays/sumOddLengthSubarrays"
	"testing"
)

func TestSumOddLengthSubarrays(t *testing.T)  {
	tests := []struct{
		arr []int
		output int
	}{
		{[]int{1,4,2,5,3}, 58},
		{[]int{1, 2}, 3},
	}

	for _, test := range tests {
		ret := sumOddLengthSubarrays.SumOddLengthSubarrays(test.arr)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.arr, test.output)
		}
	}
}