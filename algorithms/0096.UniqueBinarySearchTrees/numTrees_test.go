package main

import (
	"leetcode/algorithms/0096.UniqueBinarySearchTrees/numTrees"
	"testing"
)

func TestNumTrees(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{3, 5},
	}

	for _, test := range tests {
		ret := numTrees.NumTrees(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}