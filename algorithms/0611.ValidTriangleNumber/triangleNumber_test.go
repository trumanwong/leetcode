package main

import (
	"leetcode/algorithms/0611.ValidTriangleNumber/triangleNumber"
	"testing"
)

func TestTriangleNumber(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{2, 2, 3, 4}, 3},
	}

	for _, test := range tests {
		ret := triangleNumber.TriangleNumber(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}