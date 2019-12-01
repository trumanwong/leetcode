package zigzagLevelOrder

import . "leetcode/common/treeNode"

func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var recursive func(node *TreeNode, level int)
	recursive = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{})
		}

		if uint(level) & 1 == 1 {
			res[level] = append([]int{node.Val}, res[level]...)
		} else {
			res[level] = append(res[level], node.Val)
		}
		recursive(node.Left, level + 1)
		recursive(node.Right, level + 1)
	}
	recursive(root, 0)
	return res
}

