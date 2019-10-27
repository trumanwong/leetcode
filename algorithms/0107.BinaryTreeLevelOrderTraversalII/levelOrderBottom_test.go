package main

import (
	"leetcode/algorithms/0107.BinaryTreeLevelOrderTraversalII/levelOrderBottom"
	"testing"
)

func TestLevelOrderBottom(t *testing.T)  {
	tests := []struct{
		root   levelOrderBottom.TreeNode
		output [][]int
	}{
		{levelOrderBottom.TreeNode{3,
			&levelOrderBottom.TreeNode{9,nil,nil},
			&levelOrderBottom.TreeNode{20,
				&levelOrderBottom.TreeNode{15,nil,nil},
				&levelOrderBottom.TreeNode{7,nil,nil}}}, [][]int{
			{15,7},{9,20},{3},
		}},
	}

	for _, test := range tests {
		ret := levelOrderBottom.LevelOrderBottom(&test.root)
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
