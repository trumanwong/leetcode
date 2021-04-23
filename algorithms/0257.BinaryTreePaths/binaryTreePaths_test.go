package main

import (
	"leetcode/algorithms/0257.BinaryTreePaths/binaryTreePaths"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		root   binaryTreePaths.TreeNode
		output []string
	}{
		{binaryTreePaths.TreeNode{1,
			&binaryTreePaths.TreeNode{2, nil,
				&binaryTreePaths.TreeNode{5, nil, nil}},
			&binaryTreePaths.TreeNode{3, nil, nil}},
			[]string{"1->2->5", "1->3"}},
	}

	for _, test := range tests {
		ret := binaryTreePaths.BinaryTreePaths(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}
