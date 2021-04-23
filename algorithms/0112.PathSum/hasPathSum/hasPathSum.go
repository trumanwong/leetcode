package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum == root.Val {
		return true
	}

	return HasPathSum(root.Left, sum-root.Val) || HasPathSum(root.Right, sum-root.Val)
}
