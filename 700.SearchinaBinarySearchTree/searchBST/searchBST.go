package searchBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	res := search(root, val)
	return res
}

func search(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	L := search(root.Left, val)
	R := search(root.Right, val)
	if L != nil {
		return L
	}
	if R != nil {
		return R
	}
	return nil
}