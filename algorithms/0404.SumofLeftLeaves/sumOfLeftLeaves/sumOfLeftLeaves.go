package sumOfLeftLeaves

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L := recursive(node.Left)
		recursive(node.Right)
		res += L
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
		return 0
	}
	recursive(root)
	return res
}
