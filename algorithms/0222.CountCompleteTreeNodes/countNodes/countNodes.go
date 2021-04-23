package countNodes

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := countNodes(root.Left), countNodes(root.Right)
	return left + right + 1
}
