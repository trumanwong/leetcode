package main

import (
	"leetcode/algorithms/0203.RemoveLinkedListElements/removeElements"
	"testing"
)

func TestRemoveElements(t *testing.T)  {
	tests := []struct{
		input  removeElements.ListNode
		output *removeElements.ListNode
	}{
		{removeElements.ListNode{1,
			&removeElements.ListNode{2,
				&removeElements.ListNode{6,
					&removeElements.ListNode{2,
						&removeElements.ListNode{6, nil}}}}},
						&removeElements.ListNode{1,&removeElements.ListNode{2,&removeElements.ListNode{2, nil}}}},
	}
	for _, test := range tests {
		ret := removeElements.RemoveElements(&test.input, 6)
		for ret != nil {
			ret = ret.Next
		}
	}
}