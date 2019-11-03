package reverseKGroup

import (
	. "leetcode/common/list"
)

func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, nil}
	stack := make([]*ListNode, 0)
	p := dummy
	for {
		count := 0
		tmp := head
		for tmp != nil && count < k {
			stack = append(stack, tmp)
			tmp = tmp.Next
			count++
		}

		if count != k {
			p.Next = head
			break
		}

		// 反转操作
		for len(stack) != 0 {
			p.Next = stack[len(stack) - 1]
			p = p.Next
			stack = stack[:len(stack) - 1]
		}

		// 与剩下的链表连接
		p.Next = tmp
		head = tmp
	}
	return dummy.Next
}