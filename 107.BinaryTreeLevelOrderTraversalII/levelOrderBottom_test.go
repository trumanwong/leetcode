package main

import (
	"testing"
	. "truman.com/leetcode/107.BinaryTreeLevelOrderTraversalII/levelOrderBottom"
)

func TestLevelOrderBottom(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output [][]int
	}{
		{TreeNode{3,
			&TreeNode{9,nil,nil},
			&TreeNode{20,
				&TreeNode{15,nil,nil},
				&TreeNode{7,nil,nil}}}, [][]int{
			{15,7},{9,20},{3},
		}},
	}

	for _, test := range tests {
		ret := LevelOrderBottom(&test.root)
		for i, values := range ret {
			for j, v := range values {
				if v != test.output[i][j] {
					t.Errorf("Got %v; expected %v", ret, test.output)
					break
				}
			}
		}
	}
}
