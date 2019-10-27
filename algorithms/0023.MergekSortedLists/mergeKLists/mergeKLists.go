package mergeKLists

import "sort"

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{0, nil}
	point := head
	nodes := make([]int, 0)
	for _, list := range lists {
		for list != nil {
			nodes = append(nodes, list.Val)
			list = list.Next
		}
	}

	sort.Ints(nodes)
	for _, v := range nodes {
		point.Next = &ListNode{v, nil}
		point = point.Next
	}
	return head.Next
}