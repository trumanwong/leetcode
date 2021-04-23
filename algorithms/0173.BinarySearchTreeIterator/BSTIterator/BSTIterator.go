package BSTIterator

import . "leetcode/common/treeNode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodes []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	nodes := make([]*TreeNode, 0)
	for root != nil {
		nodes = append(nodes, root)
		root = root.Left
	}
	return BSTIterator{nodes: nodes}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	temp := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	res := temp.Val
	temp = temp.Right
	for temp != nil {
		this.nodes = append(this.nodes, temp)
		temp = temp.Left
	}
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
