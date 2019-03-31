package main

import (
	"testing"
	. "truman.com/leetcode/104.MaximumDepthofBinaryTree/maxDepth"
)

func TestMaxDepth(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output int
	}{
		{TreeNode{3,&TreeNode{9,nil,nil},
			&TreeNode{20,&TreeNode{15,nil,nil},
				&TreeNode{7,nil,nil}}},3},
	}

	for _, test := range tests {
		ret := MaxDepth(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}