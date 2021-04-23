package main

import (
	"leetcode/algorithms/0852.PeakIndexinaMountainArray/peakIndexInMountainArray"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
	}

	for _, test := range tests {
		ret := peakIndexInMountainArray.PeakIndexInMountainArray(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
