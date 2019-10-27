package main

import (
	"leetcode/algorithms/0110.BalancedBinaryTree/isBalanced"
	"testing"
)

func TestIsBalanced(t *testing.T)  {
	tests := []struct{
		root   isBalanced.TreeNode
		output bool
	}{
		{isBalanced.TreeNode{3,
			&isBalanced.TreeNode{9, nil, nil},
			&isBalanced.TreeNode{20,
				&isBalanced.TreeNode{15, nil, nil},
				&isBalanced.TreeNode{7, nil, nil}}}, true},
		{isBalanced.TreeNode{1, &isBalanced.TreeNode{2,
			&isBalanced.TreeNode{3, &isBalanced.TreeNode{4, nil, nil}, &isBalanced.TreeNode{4, nil, nil}},
			&isBalanced.TreeNode{3, nil, nil}}, &isBalanced.TreeNode{2, nil, nil}}, false},
	}

	for _, test := range tests {
		ret := isBalanced.IsBalanced(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}