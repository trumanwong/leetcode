package main

import (
	"fmt"
	"leetcode/algorithms/0095.UniqueBinarySearchTreesII/generateTrees"
	"testing"
)

func TestGenerateTrees(t *testing.T)  {
	tests := []struct{
		n int
		output [][]interface{}
	}{
		{3, [][]interface{}{{1, nil, 2, nil, 3}, {1, nil, 3, 2}, {2, 1, 3}, {3, 1, nil, nil, 2}, {3, 2, nil, 1}}},
	}

	for _, test := range tests {
		ret := generateTrees.GenerateTrees(test.n)
		fmt.Println(ret)
	}
}