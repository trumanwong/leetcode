package findTarget

import (
	. "leetcode/common/treeNode"
	"sort"
)

func FindTarget(root *TreeNode, k int) bool {
	arr := make([]int, 0)
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		recursive(node.Left)
		recursive(node.Right)
	}
	recursive(root)
	sort.Ints(arr)
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l]+arr[r] == k {
			return true
		} else if arr[l]+arr[r] > k {
			r--
		} else {
			l++
		}
	}
	return false
}
