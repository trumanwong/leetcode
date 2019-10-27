package main

import (
	"leetcode/algorithms/0485.MaxConsecutiveOnes/findMaxConsecutiveOnes"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
	}

	for _, test := range tests {
		ret := findMaxConsecutiveOnes.FindMaxConsecutiveOnes(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}