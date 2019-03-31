/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp1 := &l1
	temp2 := *temp1
	temp3 := &l2
	temp4 := *temp3
	dumpyHead := &ListNode{Val: 0}
	curr := dumpyHead
	carry, x, y, sum := 0, 0, 0, 0
	for ; temp2 != nil || temp4 != nil; {
		if temp2 != nil {
			x = temp2.Val
		} else {
			x = 0
		}
		if temp4 != nil {
			y = temp4.Val
		} else {
			y = 0
		}
		sum = carry + x + y
		carry = sum / 10
		mod := sum % 10
		curr.Next = &ListNode{Val: mod}
		curr = curr.Next
		if temp2 != nil {
			temp2 = temp2.Next
		}
		if temp4 != nil {
			temp4 = temp4.Next
		}
		if carry > 0 {
			curr.Next = &ListNode{Val: carry}
		}
	}
	return dumpyHead.Next
}
