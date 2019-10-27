package main

import (
	"leetcode/algorithms/0023.MergekSortedLists/mergeKLists"
	"testing"
)

type TestStruct struct {
	lists []*mergeKLists.ListNode
	output *mergeKLists.ListNode
}

func TestMergeKLists(t *testing.T)  {
	input := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	output := []int{1, 1, 2, 3, 4, 4, 5, 6}
	test := TestStruct{}
	for _, arr := range input {
		node := &mergeKLists.ListNode{0, nil}
		temp := node
		for _, v := range arr {
			temp.Next = &mergeKLists.ListNode{v, nil}
			temp = temp.Next
		}
		test.lists = append(test.lists, node.Next)
	}

	node := &mergeKLists.ListNode{0, nil}
	temp := node
	for _, v := range output {
		temp.Next = &mergeKLists.ListNode{v, nil}
		temp = temp.Next
	}
	test.output = node.Next

	ret := mergeKLists.MergeKLists(test.lists)
	for ret != nil && test.output != nil {
		if ret.Val != test.output.Val {
			t.Errorf("not excepted %d, excepted %d", ret.Val, test.output.Val)
			break
		}
		ret = ret.Next
		test.output = test.output.Next
	}
}