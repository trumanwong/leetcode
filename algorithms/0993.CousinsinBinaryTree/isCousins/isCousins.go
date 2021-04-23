package isCousins

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	res, m := make(map[int]int), make(map[int]int)
	var recursive func(node *TreeNode, curr int, parent int)
	recursive = func(node *TreeNode, curr int, parent int) {
		if node == nil {
			return
		}
		res[node.Val], m[node.Val] = curr, parent
		curr++
		recursive(node.Left, curr, node.Val)
		recursive(node.Right, curr, node.Val)
	}
	recursive(root, 0, -1)
	return res[x] == res[y] && m[x] != m[y]
}
