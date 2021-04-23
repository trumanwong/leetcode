package main

import (
	"leetcode/algorithms/0812.LargestTriangleAreat/largestTriangleArea"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	tests := []struct {
		points [][]int
		output float64
	}{
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}, 2},
	}

	for _, test := range tests {
		ret := largestTriangleArea.LargestTriangleArea(test.points)
		if ret != test.output {
			t.Errorf("Got %f for input %v; expected %f", ret, test.points, test.output)
		}
	}
}
