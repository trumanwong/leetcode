package binaryTreePaths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var recursivePaths func(*TreeNode, string)
	recursivePaths = func (node *TreeNode, curr string)  {
		if node == nil {
			return
		}
		if node.Right == nil && node.Left == nil {
			curr += strconv.Itoa(node.Val)
			res = append(res, curr)
			return
		} else {
			curr += strconv.Itoa(node.Val) + "->"
		}
		recursivePaths(node.Left, curr)
		recursivePaths(node.Right, curr)
	}
	recursivePaths(root, "")
	return res
}
