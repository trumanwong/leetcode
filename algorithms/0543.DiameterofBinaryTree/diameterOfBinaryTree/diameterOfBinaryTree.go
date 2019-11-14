package diameterOfBinaryTree

import . "leetcode/common/treeNode"

func DiameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := recursive(node.Left)
		right := recursive(node.Right)
		res = max(left + right, res)
		return max(left + 1, right + 1)
	}

	recursive(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}