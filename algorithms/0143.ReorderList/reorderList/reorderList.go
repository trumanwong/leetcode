package reorderList

import (
	. "leetcode/common/list"
)

func ReorderList(head *ListNode) {
	queue := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		queue = append(queue, cur)
		cur = cur.Next
	}

	for len(queue) > 0 {
		if cur == nil {
			cur = queue[0]
			queue = queue[1:]
		} else {
			cur.Next = queue[0]
			queue = queue[1:]
			cur = cur.Next
		}
		if len(queue) == 0 {
			break
		}
		cur.Next = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = nil
	}
}
