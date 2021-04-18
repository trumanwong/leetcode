package isEvenOddTree

import (
	. "leetcode/common/treeNode"
)

func IsEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 0
	res := true
	for len(queue) > 0 {
		temp := []*TreeNode{}
		for i := 0; i < len(queue) - 1; i++ {
			// 判断是否按照偶数层递增，奇数层递减
			if depth % 2 != 0 {
				if queue[i].Val <= queue[i + 1].Val {
					res = false
				}
			} else if queue[i].Val >= queue[i + 1].Val {
				res = false
			}
		}
		for _, node := range queue {
			// 遍历所有节点判断node的值和深度奇偶性是否不同
			if node.Val % 2 == depth % 2 {
				res = false
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		depth++
		queue = temp
	}
	return res
}