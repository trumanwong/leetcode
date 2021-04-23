package rangeSumBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= L && node.Val <= R {
			res += node.Val
		}
		recursive(node.Left)
		recursive(node.Right)
	}
	recursive(root)
	return res
}
