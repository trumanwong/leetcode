package reverseList

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var res *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = res
		res = curr
		curr = nextTemp
	}
	return res
}