package bstToGst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	cur := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		cur += node.Val
		node.Val = cur
		if node.Left != nil {
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}