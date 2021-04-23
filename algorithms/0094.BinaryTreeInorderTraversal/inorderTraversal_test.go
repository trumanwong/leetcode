package main

import (
	"leetcode/algorithms/0094.BinaryTreeInorderTraversal/inorderTraversal"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root   inorderTraversal.TreeNode
		output []int
	}{
		{inorderTraversal.TreeNode{1, nil, &inorderTraversal.TreeNode{
			2, &inorderTraversal.TreeNode{3, nil, nil}, nil}}, []int{1, 3, 2}},
	}

	for _, test := range tests {
		ret := inorderTraversal.InorderTraversal(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}
