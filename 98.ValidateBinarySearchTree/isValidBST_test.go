package main

import (
	"testing"
	. "truman.com/leetcode/98.ValidateBinarySearchTree/isValidBST"
)

func TestIsValidBST(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output bool
	}{
		{TreeNode{2,&TreeNode{1,nil,nil},&TreeNode{3,nil,nil}},true},
		{TreeNode{5,&TreeNode{1,nil,nil},&TreeNode{4,&TreeNode{3,nil,nil},&TreeNode{6,nil,nil}}}, false},
	}

	for _, test :=range tests {
		ret := IsValidBST(&test.root)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}