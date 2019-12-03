package buildTree


import . "leetcode/common/treeNode"

var (
	pre_idx int
	idx_map map[int]int
)

func BuildTree(preorder []int, inorder []int) *TreeNode {
	pre_idx = 0
	idx_map = make(map[int]int)
	idx := 0
	for _, v := range inorder {
		idx_map[v] = idx
		idx++
	}
	return helper(0, len(inorder), preorder, inorder)
}

func helper(left, right int, preorder []int, inorder []int) *TreeNode {
	if left == right {
		return nil
	}

	root_val := preorder[pre_idx]
	root := TreeNode{root_val, nil, nil}

	index := idx_map[root.Val]
	pre_idx++
	root.Left = helper(left, index, preorder, inorder)
	root.Right = helper(index + 1, right, preorder, inorder)
	return &root
}