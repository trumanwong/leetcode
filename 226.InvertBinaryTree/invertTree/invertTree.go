package invertTree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Right)
	invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	return root
}