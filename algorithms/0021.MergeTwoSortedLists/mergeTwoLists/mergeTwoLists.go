package mergeTwoLists

import (
	. "leetcode/common/list"
)

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = MergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = MergeTwoLists(l1.Next, l2)
	}
	return res
}