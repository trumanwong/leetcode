package main

import (
	"testing"
	. "truman.com/leetcode/687.LongestUnivaluePath/longestUnivaluePath"
)

func TestArrowLength(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output int
	}{
		{TreeNode{5,
			&TreeNode{4,
				&TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}},
			&TreeNode{5, nil, &TreeNode{5, nil, nil}}}, 2},
		{TreeNode{1,
			&TreeNode{4,
				&TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
			&TreeNode{5, nil, &TreeNode{5, nil, nil}}}, 2},
	}
	for _, test := range tests {
		ret := LongestUnivaluePath(&test.root)
		if ret != test.output {
			t.Errorf("Got %d; expected %d", ret, test.output)
		}
	}
}