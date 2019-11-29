package main

import (
	"leetcode/algorithms/0598.RangeAdditionII/maxCount"
	"testing"
)

func TestMaxCount(t *testing.T)  {
	tests := []struct{
		m int
		n int
		ops [][]int
		output int
	}{
		{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
	}

	for _, test := range tests {
		ret := maxCount.MaxCount(test.m, test.n, test.ops)
		if ret != test.output {
			t.Errorf("Got %d for input m = %d, n = %d, ops = %v; expeted %d", ret, test.m, test.n, test.ops, test.output)
		}
	}
}