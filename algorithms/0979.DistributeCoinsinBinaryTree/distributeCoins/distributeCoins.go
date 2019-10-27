package distributeCoins

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L := dfs(node.Left)
		R := dfs(node.Right)
		res += abs(L) + abs(R)
		return node.Val + L + R - 1
	}
	dfs(root)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}