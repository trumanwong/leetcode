package main

import (
	"testing"
	"truman.com/leetcode/561.ArrayPartitionI/arrayPairSum"
)

func TestArrayPairSum(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1,4,3,2},4},
	}

	for _, test := range tests {
		ret := arrayPairSum.ArrayPairSum(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}