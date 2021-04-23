package main

import (
	"leetcode/algorithms/1128.NumberofEquivalentDominoPairs/numEquivDominoPairs"
	"testing"
)

func TestNumEquivDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		output   int
	}{
		{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, 1},
	}

	for _, test := range tests {
		ret := numEquivDominoPairs.NumEquivDominoPairs(test.dominoes)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.dominoes, test.output)
		}
	}
}
