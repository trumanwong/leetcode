package swapPairs

import (
	. "leetcode/common/list"
)

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	temp := next.Next
	next.Next = head
	head.Next = SwapPairs(temp)
	return next
}