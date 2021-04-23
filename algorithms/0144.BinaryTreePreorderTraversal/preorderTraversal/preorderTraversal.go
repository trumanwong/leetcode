package preorderTraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	res := []int{}
	recursiveTreeNode(root, &res)
	return res
}

func recursiveTreeNode(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		recursiveTreeNode(root.Left, res)
		recursiveTreeNode(root.Right, res)
	}
}
