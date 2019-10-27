package reverseBetween

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	mNode, preM, nNode := head, dummy, head
	for i := 1; i < m; i++ {
		preM = mNode
		mNode = mNode.Next
	}
	for i := 1; i < n; i++ {
		nNode = nNode.Next
	}
	for mNode != nNode {
		preM.Next = mNode.Next
		mNode.Next = nNode.Next
		nNode.Next = mNode
		mNode = preM.Next
	}
	return dummy.Next
}