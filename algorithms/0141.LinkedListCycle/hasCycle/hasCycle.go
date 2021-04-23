package hasCycle

import (
	. "leetcode/common/list"
)

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return true
}
