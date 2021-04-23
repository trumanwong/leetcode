package sumRootToLeaf

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	res, mod := 0, 1000000007
	var dfs func(root *TreeNode, cur int)
	dfs = func(root *TreeNode, cur int) {
		if root.Left == nil && root.Right == nil {
			res += cur
			res %= mod
			return
		}
		if root.Left != nil {
			dfs(root.Left, (cur*2+root.Left.Val)%mod)
		}
		if root.Right != nil {
			dfs(root.Right, (cur*2+root.Right.Val)%mod)
		}
	}
	dfs(root, root.Val)
	return res
}
