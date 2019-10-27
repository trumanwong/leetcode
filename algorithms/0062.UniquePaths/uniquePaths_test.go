package main

import (
	"leetcode/algorithms/0062.UniquePaths/uniquePaths"
	"testing"
)

func TestUniquePaths(t *testing.T)  {
	tests := []struct{
		m int
		n int
		output int
	}{
		{3, 2, 3},
		{7, 3, 28},
	}

	for _, test := range tests {
		ret := uniquePaths.UniquePaths(test.m, test.n)
		if ret != test.output {
			t.Errorf("Got %d for input m = %d, n = %d; expected %d", ret, test.m, test.n, test.output)
		}
	}
}