package main

import (
	"testing"
	. "truman.com/leetcode/129.SumRoottoLeafNumbers/sumNumbers"
)

func TestSumNumbers(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output int
	}{
		{TreeNode{1,
			&TreeNode{2,nil,nil},
			&TreeNode{3,nil,nil}}, 25},
		{TreeNode{4,
			&TreeNode{9,&TreeNode{5,nil,nil},
				&TreeNode{1,nil,nil}},
				&TreeNode{0,nil,nil}},1026},
	}

	for _, test := range tests {
		ret := SumNumbers(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}