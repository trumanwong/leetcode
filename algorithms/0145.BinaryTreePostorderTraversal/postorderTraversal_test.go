package main

import (
	"testing"
	"truman.com/leetcode/algorithms/145.BinaryTreePostorderTraversal/postorderTraversal"
)

func TestPostorderTraversal(t *testing.T)  {
	tests := []struct{
		root   postorderTraversal.TreeNode
		output []int
	}{
		{postorderTraversal.TreeNode{1,nil,&postorderTraversal.TreeNode{
			2,&postorderTraversal.TreeNode{3,nil,nil},nil}}, []int{3,2,1}},
	}

	for _, test := range tests {
		ret := postorderTraversal.PostorderTraversal(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}