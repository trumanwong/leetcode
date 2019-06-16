package main

import (
	"testing"
	"truman.com/leetcode/477.TotalHammingDistance/totalHammingDistance"
)

func TestTotalHammingDistance(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{4, 14, 2}, 6},
	}

	for _, test := range tests {
		ret := totalHammingDistance.TotalHammingDistance(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}