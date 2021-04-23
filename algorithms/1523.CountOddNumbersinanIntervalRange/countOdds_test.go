package main

import (
	"leetcode/algorithms/1523.CountOddNumbersinanIntervalRange/countOdds"
	"testing"
)

func TestCountOdds(t *testing.T) {
	tests := []struct {
		low    int
		high   int
		output int
	}{
		{3, 7, 3},
		{8, 10, 1},
	}
	for _, test := range tests {
		ret := countOdds.CountOdds(test.low, test.high)
		if ret != test.output {
			t.Errorf("Got %d for input (%d, %d); expected %d", ret, test.low, test.high, test.output)
		}
	}
}
