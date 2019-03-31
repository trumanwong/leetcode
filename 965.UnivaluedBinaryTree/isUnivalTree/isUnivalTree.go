package isUnivalTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left == nil || (root.Left.Val == root.Val && IsUnivalTree(root.Left))
	right := root.Right == nil || (root.Right.Val == root.Val && IsUnivalTree(root.Right))
	return left && right
}