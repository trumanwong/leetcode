package main

import (
	"testing"
	. "truman.com/leetcode/100.SameTree/isSameTree"
)

func TestIsSameTree(t *testing.T)  {
	tests := []struct{
		p TreeNode
		q TreeNode
		output bool
	}{
		{
			TreeNode{1,&TreeNode{2,nil,nil},&TreeNode{3,nil,nil}},
			TreeNode{1,&TreeNode{2,nil,nil},&TreeNode{3,nil,nil}},
			true,
		},
		{
			TreeNode{1,&TreeNode{2,nil,nil},nil},
			TreeNode{1,nil,&TreeNode{2,nil,nil}},
			false,
		},
		{
			TreeNode{1,&TreeNode{2,nil,nil},&TreeNode{1,nil,nil}},
			TreeNode{1,&TreeNode{1,nil,nil},&TreeNode{2,nil,nil}},
			false,
		},
	}
	for i, test := range tests {
		ret := IsSameTree(&test.p, &test.q)
		if ret != test.output {
			t.Errorf("Got %t; expected %t(%d)", ret, test.output, i)
			break
		}
	}
}