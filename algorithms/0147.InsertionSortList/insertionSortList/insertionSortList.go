package insertionSortList

import . "leetcode/common/list"

func InsertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	for head != nil {
		p := dummy
		for p.Next != nil && p.Next.Val < head.Val {
			p = p.Next
		}
		temp := p.Next
		p.Next = head
		head = head.Next
		p.Next.Next = temp
	}
	return dummy.Next
}