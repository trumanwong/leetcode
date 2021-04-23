package maxPathSum

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt32
	maxSum(root, &ans)
	return ans
}

func maxSum(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := max(0, maxSum(root.Left, ans))
	r := max(0, maxSum(root.Right, ans))

	sum := l + r + root.Val
	*ans = max(*ans, sum)
	return max(l, r) + root.Val
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
