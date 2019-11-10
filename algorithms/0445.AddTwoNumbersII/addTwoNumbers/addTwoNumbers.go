package addTwoNumbers

import . "leetcode/common/list"

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a1, a2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		a1 = append(a1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		a2 = append(a2, l2.Val)
		l2 = l2.Next
	}
	i, j := len(a1) - 1, len(a2) - 1
	carry, res := 0, make([]int, 0)
	for i >= 0 && j >= 0 {
		temp := a1[i] + a2[j] + carry
		res = append(res, temp % 10)
		if temp >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		i--
		j--
	}
	for k := i; k >= 0; k-- {
		temp := a1[k] + carry
		res = append(res, temp % 10)
		if temp >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}
	for k := j; k >= 0; k-- {
		temp := a2[k] + carry
		res = append(res, temp % 10)
		if temp >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		res = append(res, 1)
	}

	heap := &ListNode{0, nil}
	dummy := heap
	for i = len(res) - 1; i >= 0; i-- {
		dummy.Next = &ListNode{res[i], nil}
		dummy = dummy.Next
	}
	return heap.Next
}