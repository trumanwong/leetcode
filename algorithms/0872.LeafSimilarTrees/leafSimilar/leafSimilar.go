package leafSimilar

import (
	. "leetcode/common/treeNode"
	"reflect"
)

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := make([]int, 0), make([]int, 0)
	dfs(root1, &leaves1)
	dfs(root2, &leaves2)
	return reflect.DeepEqual(leaves1, leaves2)
}

func dfs(node *TreeNode, leaves *[]int)  {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}
	dfs(node.Left, leaves)
	dfs(node.Right, leaves)
}