package partition

import (
	. "leetcode/common/list"
)

func Partition(head *ListNode, x int) *ListNode {
	small, big, dummy := &ListNode{0, nil}, &ListNode{0, nil}, &ListNode{0, nil}
	s, b := small, big
	t := head
	for t != nil {
		if t.Val < x {
			s.Next = &ListNode{t.Val, nil}
			s = s.Next
		} else {
			b.Next = &ListNode{t.Val, nil}
			b = b.Next
		}
		t = t.Next
	}

	small, big = small.Next, big.Next
	pHead := dummy
	for small != nil {
		pHead.Next = small
		pHead = pHead.Next
		small = small.Next
	}

	for big != nil {
		pHead.Next = big
		pHead = pHead.Next
		big = big.Next
	}
	return dummy.Next
}
