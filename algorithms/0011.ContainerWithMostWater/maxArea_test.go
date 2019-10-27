package main

import (
	"leetcode/algorithms/0011.ContainerWithMostWater/maxArea"
	"testing"
)

func TestMaxArea(t *testing.T)  {
	tests := []struct{
		height []int
		output int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
	}

	for _, test := range tests {
		ret := maxArea.MaxArea(test.height)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.height, test.output)
		}
	}
}