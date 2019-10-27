package kthSmallest

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0, 0)
	node := root
	for len(stack) != 0 || node != nil{
		if node !=nil {
			stack = append(stack, node)
			node = node.Left
		}else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k==0 {
				return node.Val
			}
			node = node.Right
		}
	}
	return 0
}