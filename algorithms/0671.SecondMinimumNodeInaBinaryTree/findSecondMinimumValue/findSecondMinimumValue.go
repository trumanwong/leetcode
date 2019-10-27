package findSecondMinimumValue

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	first, second := math.MaxInt32, math.MaxInt32
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val < first {
			first, second = node.Val, second
		} else if node.Val != first && node.Val < second {
			second = node.Val
		}
		recursive(node.Left)
		recursive(node.Right)
	}
	recursive(root)
	if second == math.MaxInt32 {
		return -1
	}
	return second
}