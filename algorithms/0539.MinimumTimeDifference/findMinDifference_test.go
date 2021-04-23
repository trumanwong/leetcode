package main

import (
	"leetcode/algorithms/0539.MinimumTimeDifference/findMinDifference"
	"testing"
)

func TestFindMinDifference(t *testing.T) {
	tests := []struct {
		timePoints []string
		output     int
	}{
		{[]string{"23:59", "00:00"}, 1},
	}

	for _, test := range tests {
		ret := findMinDifference.FindMinDifference(test.timePoints)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.timePoints, test.output)
		}
	}
}
