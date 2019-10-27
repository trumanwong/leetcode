package main

import (
	"leetcode/algorithms/0101.SymmetricTree/isSymmetric"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root1 := isSymmetric.TreeNode{1,&isSymmetric.TreeNode{
		2,&isSymmetric.TreeNode{3,nil,nil},&isSymmetric.TreeNode{4,nil,nil},
	}, &isSymmetric.TreeNode{
		2, &isSymmetric.TreeNode{4,nil,nil},&isSymmetric.TreeNode{3,nil,nil},
	}}
	root2 := isSymmetric.TreeNode{1,&isSymmetric.TreeNode{
		2,&isSymmetric.TreeNode{3,nil,nil},nil,
	}, &isSymmetric.TreeNode{
		2, &isSymmetric.TreeNode{4,nil,nil},nil,
	}}
	tests := []struct {
		root   isSymmetric.TreeNode
		output bool
	}{
		{root1,true},
		{root2,false},
	}
	for i, test := range tests {
		ret := isSymmetric.IsSymmetric(&test.root)
		if ret != test.output {
			t.Errorf("Got %t for input root%d;expected %t", ret, i + 1, test.output)
		}
	}
}