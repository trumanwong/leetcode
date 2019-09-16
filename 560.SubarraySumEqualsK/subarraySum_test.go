package main

import (
	"testing"
	"truman.com/leetcode/560.SubarraySumEqualsK/subarraySum"
)

func TestSubarraySum(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		output int
	}{
		{[]int{1,1,1},2,2},
	}

	for _, test := range tests {
		ret := subarraySum.SubarraySum(test.nums, test.k)
		if ret != test.output {
			t.Errorf("Got %d for nums = %v, k = %d; expected %d", ret, test.nums, test.k, test.output)
		}
	}
}