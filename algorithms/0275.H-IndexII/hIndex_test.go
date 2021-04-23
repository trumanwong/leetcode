package main

import (
	"leetcode/algorithms/0275.H-IndexII/hIndex"
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		output    int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
	}

	for _, test := range tests {
		ret := hIndex.HIndex(test.citations)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.citations, test.output)
		}
	}
}
