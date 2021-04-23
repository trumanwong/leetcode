package minDepth

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	minLeft := MinDepth(root.Left)
	minRight := MinDepth(root.Right)
	if minLeft < minRight {
		return minLeft + 1
	} else {
		return minRight + 1
	}
}
