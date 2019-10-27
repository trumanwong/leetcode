package main

import (
	"leetcode/algorithms/0687.LongestUnivaluePath/longestUnivaluePath"
	"testing"
)

func TestArrowLength(t *testing.T)  {
	tests := []struct{
		root   longestUnivaluePath.TreeNode
		output int
	}{
		{longestUnivaluePath.TreeNode{5,
			&longestUnivaluePath.TreeNode{4,
				&longestUnivaluePath.TreeNode{1, nil, nil}, &longestUnivaluePath.TreeNode{1, nil, nil}},
			&longestUnivaluePath.TreeNode{5, nil, &longestUnivaluePath.TreeNode{5, nil, nil}}}, 2},
		{longestUnivaluePath.TreeNode{1,
			&longestUnivaluePath.TreeNode{4,
				&longestUnivaluePath.TreeNode{4, nil, nil}, &longestUnivaluePath.TreeNode{4, nil, nil}},
			&longestUnivaluePath.TreeNode{5, nil, &longestUnivaluePath.TreeNode{5, nil, nil}}}, 2},
	}
	for _, test := range tests {
		ret := longestUnivaluePath.LongestUnivaluePath(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}