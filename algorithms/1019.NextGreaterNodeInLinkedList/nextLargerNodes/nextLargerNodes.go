package nextLargerNodes

import . "leetcode/common/list"

func NextLargerNodes(head *ListNode) []int {
	res := make([]int, 0)
	stack := make([]int, 0)
	i := 0

	for head != nil {
		for len(stack) > 0 && head.Val > res[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = head.Val
			stack = stack[:len(stack)-1]
		}

		res = append(res, head.Val)
		stack = append(stack, i)
		head = head.Next
		i++
	}

	for len(stack) > 0 {
		res[stack[0]] = 0
		stack = stack[1:]
	}
	return res
}
