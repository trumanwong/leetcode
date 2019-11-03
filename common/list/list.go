package list

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func Constructor(arr []int) *ListNode {
	res := &ListNode{0, nil}
	dummy := res
	for _, v := range arr {
		dummy.Next = &ListNode{v, nil}
		dummy = dummy.Next
	}
	return res.Next
}

func (node *ListNode) ToArray() []int {
	dummy := node
	res := make([]int, 0)
	for dummy != nil {
		res = append(res, dummy.Val)
		dummy = dummy.Next
	}
	return res
}

func (node *ListNode) String() string {
	dummy, res := node, make([]int, 0)
	for dummy != nil {
		res = append(res, dummy.Val)
		dummy = dummy.Next
	}
	return fmt.Sprintf("%v", res)
}