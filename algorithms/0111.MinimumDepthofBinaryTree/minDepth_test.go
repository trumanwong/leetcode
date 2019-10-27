package main

import (
	"leetcode/algorithms/0111.MinimumDepthofBinaryTree/minDepth"
	"testing"
)

func TestMinDepth(t *testing.T)  {
	tests := []struct{
		root   minDepth.TreeNode
		output int
	}{
		{minDepth.TreeNode{3, &minDepth.TreeNode{9, nil, nil},
			&minDepth.TreeNode{20, &minDepth.TreeNode{15, nil, nil},
				&minDepth.TreeNode{7, nil, nil}}}, 2},
		{minDepth.TreeNode{1, &minDepth.TreeNode{2, &minDepth.TreeNode{4,nil,nil}, nil},
			&minDepth.TreeNode{3, nil,
				&minDepth.TreeNode{5, nil, nil}}}, 3},
	}
	for _, test := range tests {
		ret := minDepth.MinDepth(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}