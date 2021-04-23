package postorderTraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		res = append([]int{node.Val}, res...)
	}

	return res
}

func recursiveTreeNode(root *TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			recursiveTreeNode(root.Left, res)
		}
		if root.Right != nil {
			recursiveTreeNode(root.Right, res)
		}
		*res = append(*res, root.Val)
	}
}
