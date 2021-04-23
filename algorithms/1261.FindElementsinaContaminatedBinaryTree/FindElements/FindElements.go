package FindElements

import (
	. "leetcode/common/treeNode"
)

type FindElements struct {
	root *TreeNode
	nums map[int]int
}

func FindElementsConstructor(root *TreeNode) FindElements {
	root.Val = 0
	nums := make(map[int]int)
	nums[0]++
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		temp := make([]*TreeNode, len(nodes))
		copy(temp, nodes)
		nodes = make([]*TreeNode, 0)
		for _, node := range temp {
			if node.Left != nil {
				node.Left.Val = 2*node.Val + 1
				nums[node.Left.Val]++
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 2
				nums[node.Right.Val]++
				nodes = append(nodes, node.Right)
			}
		}
	}
	return FindElements{root: root, nums: nums}
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.nums[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
