package main

import (
	"leetcode/algorithms/0997.FindtheTownJudge/findJudge"
	"testing"
)

func TestFindJudge(t *testing.T)  {
	tests := []struct{
		N int
		trust [][]int
		output int
	}{
		{2, [][]int{{1, 2}}, 2},
		{3, [][]int{{1, 3}, {2, 3}}, 3},
		{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
		{3, [][]int{{1, 2}, {2, 3}}, -1},
	}

	for _, test := range tests {
		ret := findJudge.FindJudge(test.N, test.trust)
		if ret != test.output {
			t.Errorf("Got %d for input N = %d, trust = %v; expected %d", ret, test.N, test.trust, test.output)
		}
	}
}