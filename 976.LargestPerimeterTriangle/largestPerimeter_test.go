package main

import (
	"testing"
	"truman.com/leetcode/976.LargestPerimeterTriangle/largestPerimeter"
)

func TestLargestPerimeter(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{2,1,2},5},
		{[]int{1,2,1},0},
		{[]int{3,2,3,4},10},
		{[]int{3,6,2,3},8},
	}

	for _, test := range tests {
		ret := largestPerimeter.LargestPerimeter(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
