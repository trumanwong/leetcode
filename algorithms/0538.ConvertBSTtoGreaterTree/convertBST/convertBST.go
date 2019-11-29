package convertBST

import (
	. "leetcode/common/treeNode"
)

var sum int
func ConvertBST(root *TreeNode) *TreeNode {
	sum = 0
	//recursive(root)

	node := root
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}
		node = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		sum += node.Val
		node.Val = sum

		node = node.Left
	}
	return root
}

func recursive(root *TreeNode)  {
	if root == nil {
		return
	}
	recursive(root.Right)
	sum += root.Val
	root.Val = sum
	recursive(root.Left)
}