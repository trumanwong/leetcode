package deleteDuplicates

import (
	. "leetcode/common/list"
)

func DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
