package treeNode

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Constructor(start int, nums []interface{}) *TreeNode {
	if nums[start] == nil {
		return nil
	}

	node := TreeNode{nums[start].(int), nil, nil}
	l, r := 2 * start + 1, 2 * start + 2
	if l > len(nums) - 1 {
		node.Left = nil
	} else {
		node.Left = Constructor(l, nums)
	}

	if r > len(nums) - 1 {
		node.Right = nil
	} else {
		node.Right = Constructor(r, nums)
	}
	return &node
}

// 前序遍历 根节点->左子树->右子树
func (this *TreeNode) PreOrderTraversal() (res []int) {
	stack := make([]*TreeNode, 0)
	pNode := this
	for pNode != nil || len(stack) > 0 {
		if pNode != nil {
			res = append(res, pNode.Val)
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			pNode = stack[len(stack) - 1].Right
			stack = stack[:len(stack) - 1]
		}
	}
	return
}

// 中序遍历 左子树->根节点->右子树
func (this *TreeNode) InOrderTraversal() (res []int) {
	stack := make([]*TreeNode, 0)
	pNode := this
	for pNode != nil || len(stack) > 0 {
		if pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			res = append(res, stack[len(stack) - 1].Val)
			pNode = stack[len(stack) - 1].Right
			stack = stack[:len(stack) - 1]
		}
	}
	return
}

// 后序遍历 左子树->右子树->根节点
func (this *TreeNode)  PostOrderTraversal() (res []int) {
	stack := []*TreeNode{this}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack)- 1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		res = append([]int{node.Val}, res...)
	}
	return
}

func (this *TreeNode) LevelOrder() (res [][]int)  {
	nodes := []*TreeNode{this}
	for len(nodes) > 0 {
		currLevel := make([]int, 0)
		temp := make([]*TreeNode, len(nodes))
		copy(temp, nodes)
		nodes = make([]*TreeNode, 0)
		for _, node := range temp {
			currLevel = append(currLevel, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, currLevel)
	}

	return
}

func (this *TreeNode) String() string {
	return fmt.Sprintf("%v", this.LevelOrder())
}