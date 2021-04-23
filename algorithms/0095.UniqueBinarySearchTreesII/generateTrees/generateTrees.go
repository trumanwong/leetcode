package generateTrees

import (
	. "leetcode/common/treeNode"
)

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTree(1, n)
}

func generateTree(start, end int) []*TreeNode {
	all_trees := make([]*TreeNode, 0)
	if start > end {
		all_trees = append(all_trees, nil)
		return all_trees
	}

	for i := start; i <= end; i++ {
		left_trees := generateTree(start, i-1)
		right_trees := generateTree(i+1, end)

		for _, l := range left_trees {
			for _, r := range right_trees {
				current_trees := TreeNode{i, l, r}
				all_trees = append(all_trees, &current_trees)
			}
		}
	}
	return all_trees
}
