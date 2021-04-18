package main

import (
	"testing"
	"leetcode/algorithms/0104.MaximumDepthofBinaryTree/maxDepth"
)

func TestMaxDepth(t *testing.T)  {
	tests := []struct{
		root   maxDepth.TreeNode
		output int
	}{
		{maxDepth.TreeNode{3,&maxDepth.TreeNode{9,nil,nil},
			&maxDepth.TreeNode{20,&maxDepth.TreeNode{15,nil,nil},
				&maxDepth.TreeNode{7,nil,nil}}}, 3},
	}

	for _, test := range tests {
		ret := maxDepth.MaxDepth(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}