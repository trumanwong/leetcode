package main

import (
	"leetcode/algorithms/1086.highFive/highFive"
	"testing"
)

func TestHighFive(t *testing.T) {
	tests := []struct {
		input  [][]int
		output [][]int
	}{
		{[][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}},
			[][]int{{1, 87}, {2, 88}}},
	}

	for _, test := range tests {
		ret := highFive.HighFive(test.input)

		for i, v := range ret {
			if v[0] != test.output[i][0] || v[1] != test.output[i][1] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}
