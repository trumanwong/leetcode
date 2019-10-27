package rotateRight

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	i := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		i++
	}
	step := k % i
	fast := head
	for i := 0; i < step; i++ {
		if fast == nil {
			break
		}
		fast = fast.Next
	}
	if fast == nil {
		return head
	}

	slow := head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	fast.Next = head
	new_head := slow.Next
	slow.Next = nil
	return new_head
}