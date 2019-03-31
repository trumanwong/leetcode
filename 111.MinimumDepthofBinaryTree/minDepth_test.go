package main

import (
	"testing"
	. "truman.com/leetcode/111.MinimumDepthofBinaryTree/minDepth"
)

func TestMinDepth(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output int
	}{
		{TreeNode{3, &TreeNode{9, nil, nil},
			&TreeNode{20, &TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}, 2},
		{TreeNode{1, &TreeNode{2, &TreeNode{4,nil,nil}, nil},
			&TreeNode{3, nil,
				&TreeNode{5, nil, nil}}}, 3},
	}
	for _, test := range tests {
		ret := MinDepth(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}