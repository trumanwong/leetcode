package levelOrder

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []int{}
		nodes := []*TreeNode{}
		for _, v := range queue {
			temp = append(temp, v.Val)
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}
		res = append(res, temp)
		queue = nodes
	}
	return res
}
