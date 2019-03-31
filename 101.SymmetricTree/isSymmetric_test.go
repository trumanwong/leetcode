package main

import (
	"testing"
	. "truman.com/leetcode/101.SymmetricTree/isSymmetric"
)

func TestIsSymmetric(t *testing.T) {
	root1 := TreeNode{1,&TreeNode{
		2,&TreeNode{3,nil,nil},&TreeNode{4,nil,nil},
	}, &TreeNode{
		2, &TreeNode{4,nil,nil},&TreeNode{3,nil,nil},
	}}
	root2 := TreeNode{1,&TreeNode{
		2,&TreeNode{3,nil,nil},nil,
	}, &TreeNode{
		2, &TreeNode{4,nil,nil},nil,
	}}
	tests := []struct {
		root TreeNode
		output bool
	}{
		{root1,true},
		{root2,false},
	}
	for i, test := range tests {
		ret := IsSymmetric(&test.root)
		if ret != test.output {
			t.Errorf("Got %t for input root%d;expected %t", ret, i + 1, test.output)
		}
	}
}