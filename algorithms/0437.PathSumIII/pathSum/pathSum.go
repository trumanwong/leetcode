package pathSum

import . "leetcode/common/treeNode"

func PathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return PathSum(root.Left, sum) + PathSum(root.Right, sum) + dfs(root, sum)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val == sum {
		count = 1
	}
	return count + dfs(node.Left, sum - node.Val) + dfs(node.Right, sum - node.Val)
}