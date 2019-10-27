package inorderTraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	recursiveTreeNodeLeft(root.Left, &res)
	res = append(res, root.Val)
	recursiveTreeNodeRight(root.Right, &res)
	return res
}

func recursiveTreeNodeLeft(root *TreeNode, res *[]int) {
	if root != nil {
		recursiveTreeNodeLeft(root.Left, res)
		*res = append(*res, root.Val)
		recursiveTreeNodeRight(root.Right, res)
	}
}

func recursiveTreeNodeRight(root *TreeNode, res *[]int) {
	if root != nil {
		recursiveTreeNodeLeft(root.Left, res)
		*res = append(*res, root.Val)
		recursiveTreeNodeLeft(root.Right, res)
	}
}