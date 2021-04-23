package main

import (
	"leetcode/algorithms/0098.ValidateBinarySearchTree/isValidBST"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root   isValidBST.TreeNode
		output bool
	}{
		{isValidBST.TreeNode{2, &isValidBST.TreeNode{1, nil, nil}, &isValidBST.TreeNode{3, nil, nil}}, true},
		{isValidBST.TreeNode{5, &isValidBST.TreeNode{1, nil, nil}, &isValidBST.TreeNode{4, &isValidBST.TreeNode{3, nil, nil}, &isValidBST.TreeNode{6, nil, nil}}}, false},
	}

	for _, test := range tests {
		ret := isValidBST.IsValidBST(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}
