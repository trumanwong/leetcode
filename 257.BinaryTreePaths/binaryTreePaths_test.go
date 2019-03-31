package main

import (
	"testing"
	. "truman.com/leetcode/257.BinaryTreePaths/binaryTreePaths"
)

func TestBinaryTreePaths(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output []string
	}{
		{TreeNode{1,
			&TreeNode{2,nil,
				&TreeNode{5,nil,nil}},
				&TreeNode{3,nil,nil}},
				[]string{"1->2->5","1->3"}},
	}

	for _, test := range tests {
		ret := BinaryTreePaths(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}