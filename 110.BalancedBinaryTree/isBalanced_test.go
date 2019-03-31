package main

import (
	"testing"
	. "truman.com/leetcode/110.BalancedBinaryTree/isBalanced"
)

func TestIsBalanced(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output bool
	}{
		{TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}, true},
		{TreeNode{1, &TreeNode{2,
			&TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
			&TreeNode{3, nil, nil}}, &TreeNode{2, nil, nil}}, false},
	}

	for _, test := range tests {
		ret := IsBalanced(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}