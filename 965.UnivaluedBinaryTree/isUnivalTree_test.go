package main

import (
	"testing"
	. "truman.com/leetcode/965.UnivaluedBinaryTree/isUnivalTree"
)

func TestIsUnivalTree(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output bool
	}{
		{TreeNode{1,
			&TreeNode{1,&TreeNode{1,nil,nil},
				&TreeNode{1,nil,nil}},
				&TreeNode{1,nil,&TreeNode{1,nil,nil}}},true},
		{TreeNode{2,
			&TreeNode{2,
				&TreeNode{5,nil,nil},
				&TreeNode{2,nil,nil}},
				&TreeNode{2,nil,nil}},false},
	}
	for _, test := range tests {
		ret := IsUnivalTree(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}