package deleteDuplicates

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, nil}
	preNode, currNode, realNode := dummy, head, dummy
	for currNode != nil {
		if (preNode == dummy || preNode.Val != currNode.Val) && (currNode.Next == nil || currNode.Val != currNode.Next.Val) {
			realNode.Next = currNode
			realNode = currNode
		}

		preNode = currNode
		currNode = currNode.Next
		preNode.Next = nil
	}
	return dummy.Next
}