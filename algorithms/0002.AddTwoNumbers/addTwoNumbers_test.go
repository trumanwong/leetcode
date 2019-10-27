package main

import (
	"leetcode/algorithms/0002.AddTwoNumbers/addTwoNumbers"
	"testing"
)

type TestStruct struct {
	l1     *addTwoNumbers.ListNode
	l2     *addTwoNumbers.ListNode
	output *addTwoNumbers.ListNode
}

func TestAddTwoNumbers(t *testing.T)  {
	arr1, arr2, output := []int{2,4,3}, []int{5,6,4}, []int{7,0,8}
	l1, l2 := &addTwoNumbers.ListNode{arr1[0], nil}, &addTwoNumbers.ListNode{arr2[0], nil}
	for i := 1; i < len(arr1); i++ {
		l1.Next = &addTwoNumbers.ListNode{arr1[i], nil}
	}
	for i := 1; i < len(arr2); i++ {
		l2.Next = &addTwoNumbers.ListNode{arr2[i], nil}
	}
	ret := addTwoNumbers.AddTwoNumbers(l1, l2)
	i := 0
	for ret != nil {
		if ret.Val != output[i] {
			t.Errorf("error occured; expected %v", output)
			break
		}
		ret = ret.Next
	}
}