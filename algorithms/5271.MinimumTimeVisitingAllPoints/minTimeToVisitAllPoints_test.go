package main

import (
	"leetcode/algorithms/5271.MinimumTimeVisitingAllPoints/minTimeToVisitAllPoints"
	"testing"
)

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		output int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
	}

	for _, test := range tests {
		ret := minTimeToVisitAllPoints.MinTimeToVisitAllPoints(test.points)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.points, test.output)
		}
	}
}
