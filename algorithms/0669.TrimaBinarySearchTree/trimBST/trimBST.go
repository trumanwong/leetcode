package trimBST

import . "leetcode/common/treeNode"

func TrimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < L {
		return TrimBST(root.Right, L, R) // 返回修剪过的右子树
	}
	if root.Val > R {
		return TrimBST(root.Left, L, R) // 返回修剪过的左子树
	}

	root.Left = TrimBST(root.Left, L, R)
	root.Right = TrimBST(root.Right, L, R)
	return root
}
