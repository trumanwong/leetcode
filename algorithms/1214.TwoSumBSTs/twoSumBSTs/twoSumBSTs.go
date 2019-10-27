package twoSumBSTs

import "sort"

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TwoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	arr1, arr2 := make([]int, 0), make([]int, 0)
	var recursive func(node *TreeNode, i int)
	hasZero := false
	recursive = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		if i == 1 {
			arr1 = append(arr1, node.Val)
		} else {
			if node.Val == 0 {
				hasZero = true
			}
			arr2 = append(arr2, node.Val)
		}
		recursive(node.Left, i)
		recursive(node.Right, i)
	}
	recursive(root1, 1)
	recursive(root2, 2)
	sort.Ints(arr1)
	sort.Ints(arr2)
	if arr1[len(arr1) - 1] + arr2[len(arr2) - 1] < target || arr1[0] + arr2[0] > target {
		return false
	}

	index := -1
	for i := 0; i < len(arr2); i++ {
		if arr2[i] > 0 {
			index = i
			break
		}
	}

	for i := len(arr1) - 1; i >= 0; i-- {
		if arr1[i] > target {
			for j := 0; arr2[j] < 0; j++ {
				if arr1[i] + arr2[j] == target {
					return true
				}
			}
		} else if arr1[i] == target && hasZero {
			return true
		} else {
			if index == -1 {
				return false
			}
			for j := index; j < len(arr2); j++ {
				if arr1[i] + arr2[j] == target {
					return true
				} else if arr1[i] + arr2[j] > target {
					break
				}
			}
		}
	}
	return false
}