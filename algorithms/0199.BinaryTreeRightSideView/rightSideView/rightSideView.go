package rightSideView

import . "leetcode/common/treeNode"

var res []int

func RightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root, 1)
	return res
}

func dfs(node *TreeNode, k int) {
	if node == nil {
		return
	}
	if k > len(res) {
		res = append(res, node.Val)
	}
	dfs(node.Right, k+1)
	dfs(node.Left, k+1)
}
