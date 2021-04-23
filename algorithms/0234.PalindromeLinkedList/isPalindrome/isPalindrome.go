package isPalindrome

import (
	. "leetcode/common/list"
)

func IsPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
