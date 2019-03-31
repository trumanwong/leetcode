package main

import (
	"fmt"
	"testing"
	. "truman.com/leetcode/113.PathSumII/pathSum"
)

func TestPathSum(t *testing.T)  {
	root := TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,&TreeNode{7,nil,nil},
				&TreeNode{2,nil,nil}},nil},
		&TreeNode{8,&TreeNode{13,nil,nil},
			&TreeNode{4,&TreeNode{5,nil,nil},&TreeNode{1,nil,nil}}}}
	tests := []struct {
		root TreeNode
		sum int
		output [][]int
	}{
		{root, 22,[][]int{{5,4,11,2},{5,8,4,5}}},
	}

	for _, test := range tests {
		ret := PathSum(&test.root, test.sum)
		fmt.Println(ret)
	}
}