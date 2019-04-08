package findMode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m, res, max := make(map[int]int), []int{}, 0
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		m[node.Val]++
		if m[node.Val] > max {
			max = m[node.Val]
		}
		recursive(node.Left)
		recursive(node.Right)
	}
	recursive(root)
	for i, v := range m {
		if v == max {
			res = append(res, i)
		}
	}
	return res
}