package removeElements

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = RemoveElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}