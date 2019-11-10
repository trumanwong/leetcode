package sortList

import (
	. "leetcode/common/list"
	"sort"
)

func SortList(head *ListNode) *ListNode {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := &ListNode{0, nil}
	heap := res
	sort.Ints(arr)
	for _, v := range arr {
		heap.Next = &ListNode{v, nil}
		heap = heap.Next
	}
	return res.Next
}