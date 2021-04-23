package main

import (
	"leetcode/algorithms/0461.HammingDistance/hammingDistance"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		output int
	}{
		{1, 4, 2},
	}

	for _, test := range tests {
		ret := hammingDistance.HammingDistance(test.x, test.y)
		if ret != test.output {
			t.Errorf("Got %d for x=%d, y=%d; expected %d", ret, test.x, test.y, test.output)
		}
	}
}
