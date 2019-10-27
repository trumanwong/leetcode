package averageOfLevels

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodes := []*TreeNode{}
		sum := 0
		count := 0
		for _, v := range queue {
			sum += v.Val
			count++
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}
		queue = nodes
		average := float64(sum) / float64(count)
		res = append(res, average)
	}

	return res
}