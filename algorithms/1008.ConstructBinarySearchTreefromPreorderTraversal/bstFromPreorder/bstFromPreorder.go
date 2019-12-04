package bstFromPreorder

import (
	. "leetcode/common/treeNode"
	"math"
)

var idx int
func BstFromPreorder(preorder []int) *TreeNode {
	idx = 0
	return helper(math.MinInt32, math.MaxInt32, preorder)
}

func helper(lower, upper int, preorder []int) *TreeNode {
	if idx == len(preorder) {
		return nil
	}

	val := preorder[idx]
	if val < lower || val > upper {
		return nil
	}

	idx++
	root := TreeNode{val, nil, nil}
	root.Left = helper(lower, val, preorder)
	root.Right = helper(val, upper, preorder)
	return &root
}