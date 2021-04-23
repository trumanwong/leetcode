package main

import (
	"leetcode/algorithms/1582.SpecialPositionsinaBinaryMatrix/numSpecial"
	"testing"
)

func TestNumSpecial(t *testing.T) {
	tests := []struct {
		mat    [][]int
		output int
	}{
		{[][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}, 1},
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for _, test := range tests {
		ret := numSpecial.NumSpecial(test.mat)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.mat, test.output)
		}
	}
}
