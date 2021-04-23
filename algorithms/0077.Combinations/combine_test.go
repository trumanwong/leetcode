package main

import (
	"leetcode/algorithms/0077.Combinations/combine"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		output [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4}}},
	}

	for _, test := range tests {
		ret := combine.Combine(test.n, test.k)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input n = %d, k = %d; expected %v", ret, test.n, test.k, test.output)
				break
			}
		}
	}
}
