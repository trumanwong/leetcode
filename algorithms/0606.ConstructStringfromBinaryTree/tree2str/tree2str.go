package tree2str

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	res := strconv.Itoa(t.Val)
	if t.Left != nil {
		res += "(" + tree2str(t.Left) + ")"
	} else {
		res += "()"
	}

	if t.Right != nil {
		res += "(" + tree2str(t.Right) + ")"
	}
	return res
}
