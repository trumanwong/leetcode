package main

import (
	"testing"
	. "truman.com/leetcode/94.BinaryTreeInorderTraversal/inorderTraversal"
)

func TestInorderTraversal(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output []int
	}{
		{TreeNode{1,nil,&TreeNode{
			2,&TreeNode{3,nil,nil},nil}}, []int{1,3,2}},
	}

	for _, test := range tests {
		ret := InorderTraversal(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}