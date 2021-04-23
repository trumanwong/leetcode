package main

import (
	"leetcode/algorithms/0100.SameTree/isSameTree"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p      isSameTree.TreeNode
		q      isSameTree.TreeNode
		output bool
	}{
		{
			isSameTree.TreeNode{1, &isSameTree.TreeNode{2, nil, nil}, &isSameTree.TreeNode{3, nil, nil}},
			isSameTree.TreeNode{1, &isSameTree.TreeNode{2, nil, nil}, &isSameTree.TreeNode{3, nil, nil}},
			true,
		},
		{
			isSameTree.TreeNode{1, &isSameTree.TreeNode{2, nil, nil}, nil},
			isSameTree.TreeNode{1, nil, &isSameTree.TreeNode{2, nil, nil}},
			false,
		},
		{
			isSameTree.TreeNode{1, &isSameTree.TreeNode{2, nil, nil}, &isSameTree.TreeNode{1, nil, nil}},
			isSameTree.TreeNode{1, &isSameTree.TreeNode{1, nil, nil}, &isSameTree.TreeNode{2, nil, nil}},
			false,
		},
	}
	for i, test := range tests {
		ret := isSameTree.IsSameTree(&test.p, &test.q)
		if ret != test.output {
			t.Errorf("Got %t; expected %t(%d)", ret, test.output, i)
			break
		}
	}
}
