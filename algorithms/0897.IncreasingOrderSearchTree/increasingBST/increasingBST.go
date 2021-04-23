package increasingBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	res := &TreeNode{0, nil, nil}
	cur := res
	for _, v := range nums {
		cur.Right = &TreeNode{v, nil, nil}
		cur = cur.Right
	}
	return res.Right
}
