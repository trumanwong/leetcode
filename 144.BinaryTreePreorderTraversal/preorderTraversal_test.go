package main

import (
	"testing"
	. "truman.com/leetcode/144.BinaryTreePreorderTraversal/preorderTraversal"
)

func TestPreorderTraversal(t *testing.T)  {
	root := TreeNode{1,nil,&TreeNode{2,&TreeNode{3,nil,nil},nil}}
	tests := []struct{
		input TreeNode
		output []int
	}{
		{root,[]int{1,2,3}},
	}
	for _, test := range tests {
		ret := PreorderTraversal(&test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break;
			}
		}
	}
}