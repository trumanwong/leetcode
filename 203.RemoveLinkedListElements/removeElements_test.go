package main

import (
	"fmt"
	"testing"
	. "truman.com/leetcode/203.RemoveLinkedListElements/removeElements"
)

func TestRemoveElements(t *testing.T)  {
	tests := []struct{
		input ListNode
		output *ListNode
	}{
		{ListNode{1,
			&ListNode{2,
				&ListNode{6,
					&ListNode{2,
						&ListNode{6, nil}}}}},
						&ListNode{1,&ListNode{2,&ListNode{2, nil}}}},
	}
	for _, test := range tests {
		ret := RemoveElements(&test.input, 6)
		for ret != nil {
			fmt.Println(ret.Val)
			ret = ret.Next
		}
	}
}