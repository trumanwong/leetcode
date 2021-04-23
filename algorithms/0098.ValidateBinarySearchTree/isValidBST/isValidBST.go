package isValidBST

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	var last int    // 用来存储中序遍历时的上一个点
	isFirst := true // 用来表示是否为中序遍历获取到的第一个点

	//中序遍历二叉树，看其是否递增
	return inorderTraversal(root, &last, &isFirst)
}

/**
 * 中序遍历二叉树，看其是否递增
 */
func inorderTraversal(root *TreeNode, last *int, isFirst *bool) bool {
	// 空，返回true
	if root == nil {
		return true
	}

	/* 中序遍历左子树 */
	// 如果左子树中序遍历序列不递增，返回false
	if !inorderTraversal(root.Left, last, isFirst) {
		return false
	}

	/* 遍历当前节点 */
	if *isFirst { // 如果当前节点是中序遍历获得的第一个节点
		*isFirst = false // 修改标志位
		*last = root.Val // 更新中序遍历过程中上一个点的值
	} else { // 如果当前节点不是中序遍历获得的第一个节点
		if *last >= root.Val { // 判断当前节点是否大于遍历的上一个点
			return false // 当前节点不大于上一个点，返回false
		} else { // 当前节点大于上一个点，更新中序遍历过程中上一个点的值
			*last = root.Val
		}
	}

	/* 中序遍历左子树 */
	// 如果右子树中序遍历序列不递增，返回false
	if !inorderTraversal(root.Right, last, isFirst) {
		return false
	}

	// 执行至此，说明中序遍历序列递增，返回true
	return true
}
