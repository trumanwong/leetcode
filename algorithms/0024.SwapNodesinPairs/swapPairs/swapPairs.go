package swapPairs

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

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