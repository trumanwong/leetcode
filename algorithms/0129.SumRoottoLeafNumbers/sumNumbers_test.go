package main

import (
	"leetcode/algorithms/0129.SumRoottoLeafNumbers/sumNumbers"
	"testing"
)

func TestSumNumbers(t *testing.T)  {
	tests := []struct{
		root   sumNumbers.TreeNode
		output int
	}{
		{sumNumbers.TreeNode{1,
			&sumNumbers.TreeNode{2,nil,nil},
			&sumNumbers.TreeNode{3,nil,nil}}, 25},
		{sumNumbers.TreeNode{4,
			&sumNumbers.TreeNode{9,&sumNumbers.TreeNode{5,nil,nil},
				&sumNumbers.TreeNode{1,nil,nil}},
				&sumNumbers.TreeNode{0,nil,nil}}, 1026},
	}

	for _, test := range tests {
		ret := sumNumbers.SumNumbers(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}