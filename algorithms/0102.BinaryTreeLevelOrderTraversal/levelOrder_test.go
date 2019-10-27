package main

import (
	"leetcode/algorithms/0102.BinaryTreeLevelOrderTraversal/levelOrder"
	"testing"
)

func TestLevelOrder(t *testing.T)  {
	tests := []struct{
		root   levelOrder.TreeNode
		output [][]int
	}{
		{levelOrder.TreeNode{3,
			&levelOrder.TreeNode{9,nil,nil},
			&levelOrder.TreeNode{20,
				&levelOrder.TreeNode{15,nil,nil},
				&levelOrder.TreeNode{7,nil,nil}}}, [][]int{
					{3},{9,20},{15,7},
		}},
	}


	for _, test := range tests {
		ret := levelOrder.LevelOrder(&test.root)
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
