package deleteNode

import (
	. "leetcode/common/list"
)

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
