package main

import (
	"leetcode/algorithms/0965.UnivaluedBinaryTree/isUnivalTree"
	"testing"
)

func TestIsUnivalTree(t *testing.T) {
	tests := []struct {
		root   isUnivalTree.TreeNode
		output bool
	}{
		{isUnivalTree.TreeNode{1,
			&isUnivalTree.TreeNode{1, &isUnivalTree.TreeNode{1, nil, nil},
				&isUnivalTree.TreeNode{1, nil, nil}},
			&isUnivalTree.TreeNode{1, nil, &isUnivalTree.TreeNode{1, nil, nil}}}, true},
		{isUnivalTree.TreeNode{2,
			&isUnivalTree.TreeNode{2,
				&isUnivalTree.TreeNode{5, nil, nil},
				&isUnivalTree.TreeNode{2, nil, nil}},
			&isUnivalTree.TreeNode{2, nil, nil}}, false},
	}
	for _, test := range tests {
		ret := isUnivalTree.IsUnivalTree(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}
