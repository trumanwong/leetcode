package main

import (
	"fmt"
	"leetcode/algorithms/0113.PathSumII/pathSum"
	"testing"
)

func TestPathSum(t *testing.T)  {
	root := pathSum.TreeNode{5,
		&pathSum.TreeNode{4,
			&pathSum.TreeNode{11,&pathSum.TreeNode{7,nil,nil},
				&pathSum.TreeNode{2,nil,nil}}, nil},
		&pathSum.TreeNode{8,&pathSum.TreeNode{13,nil,nil},
			&pathSum.TreeNode{4,&pathSum.TreeNode{5,nil,nil},&pathSum.TreeNode{1,nil,nil}}}}
	tests := []struct {
		root   pathSum.TreeNode
		sum    int
		output [][]int
	}{
		{root, 22,[][]int{{5,4,11,2},{5,8,4,5}}},
	}

	for _, test := range tests {
		ret := pathSum.PathSum(&test.root, test.sum)
		fmt.Println(ret)
	}
}