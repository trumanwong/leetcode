package levelOrderBottom

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nodes := []*TreeNode{}
		temp := []int{}
		for _, v := range stack {
			temp = append(temp, v.Val)
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}
		stack = nodes
		res = append(res, temp)
	}
	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i],res[i]
	}
	return res
}