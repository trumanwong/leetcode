package main

import (
	"leetcode/algorithms/0144.BinaryTreePreorderTraversal/preorderTraversal"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := preorderTraversal.TreeNode{1, nil, &preorderTraversal.TreeNode{2, &preorderTraversal.TreeNode{3, nil, nil}, nil}}
	tests := []struct {
		input  preorderTraversal.TreeNode
		output []int
	}{
		{root, []int{1, 2, 3}},
	}
	for _, test := range tests {
		ret := preorderTraversal.PreorderTraversal(&test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}
