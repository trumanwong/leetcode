package getMinimumDifference

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return math.MaxInt32
	}
	val, left, right := root.Val, math.MaxInt32, math.MaxInt32
	if root.Left != nil {
		l := root.Left
		for l.Right != nil {
			l = l.Right
		}
		left = val - l.Val
	}

	if root.Right != nil {
		r := root.Right
		for r.Left != nil {
			r = r.Left
		}
		right = r.Val - val
	}
	return min(min(left, right), min(getMinimumDifference(root.Left), getMinimumDifference(root.Right)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
