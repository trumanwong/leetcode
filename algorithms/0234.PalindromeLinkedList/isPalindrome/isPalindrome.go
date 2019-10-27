package isPalindrome

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}


func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l, r := 0, len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}