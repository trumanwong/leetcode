package maxLevelSum

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	ret, max, arr := 0, 0, make([]int, 10000)
	maxIndex := 0
	var recursive func(node *TreeNode, index int)
	recursive = func(node *TreeNode, index int) {
		if node == nil {
			return
		}
		if index > maxIndex {
			maxIndex = index
		}
		arr[index] += node.Val
		recursive(node.Left, index + 1)
		recursive(node.Right, index + 1)
	}
	recursive(root, 0)
	for i := 0; i < maxIndex; i++ {
		if arr[i] > max {
			max = arr[i]
			ret = i + 1
		}
	}
	return ret
}