package main

import (
	"testing"
	. "truman.com/leetcode/145.BinaryTreePostorderTraversal/postorderTraversal"
)

func TestPostorderTraversal(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output []int
	}{
		{TreeNode{1,nil,&TreeNode{
			2,&TreeNode{3,nil,nil},nil}},[]int{3,2,1}},
	}

	for _, test := range tests {
		ret := PostorderTraversal(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}