package minDiffInBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res, pre := 1 << 31 - 1, -1 << 31
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			recursive(node.Left)
		}
		res = min(res, node.Val - pre)
		pre = node.Val
		if node.Right != nil {
			recursive(node.Right)
		}
	}
	recursive(root)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}