package oddEvenList

import (
	. "leetcode/common/list"
)

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 双指针
	o, e := head, head.Next
	p1, p2 := o, e
	for p2 != nil && p2.Next != nil {
		p1.Next = p1.Next.Next
		p2.Next = p2.Next.Next
		p1, p2 = p1.Next, p2.Next
	}
	// 拼接
	p1.Next = e
	return head
}