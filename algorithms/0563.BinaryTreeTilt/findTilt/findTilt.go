package findTilt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res := 0
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := traverse(node.Left)
		right := traverse(node.Right)
		res += abs(left - right)
		return left + right + node.Val
	}
	traverse(root)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
