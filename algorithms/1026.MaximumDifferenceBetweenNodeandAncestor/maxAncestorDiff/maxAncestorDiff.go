package maxAncestorDiff

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	var recursive func(node *TreeNode, maxVal int, minVal int)
	res := 0
	recursive = func(node *TreeNode, maxVal int, minVal int) {
		if node == nil {
			return
		}
		diff1, diff2 := abs(maxVal-node.Val), abs(minVal-node.Val)
		res = max(res, max(diff1, diff2))
		maxVal = max(maxVal, node.Val)
		minVal = min(minVal, node.Val)
		recursive(node.Left, maxVal, minVal)
		recursive(node.Right, maxVal, minVal)
	}
	recursive(root, root.Val, root.Val)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
