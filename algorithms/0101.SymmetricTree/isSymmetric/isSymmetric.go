package isSymmetric

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}
