package main

import (
	"leetcode/algorithms/0707.DesignLinkedList/MyLinkedList"
	"testing"
)

func TestMyLinkedList(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output []interface{}
	}{
		{[]string{"MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"},
		[][]int{{},{1},{3},{1,2},{1},{1},{1}},
		[]interface{}{nil,nil,nil,nil,2,nil,3}},
	}

	for _, test := range tests {
		myLinkedList := MyLinkedList.MyLinkedList{}
		for i, operate := range test.operates {
			if operate == "MyLinkedList" {
				myLinkedList = MyLinkedList.Constructor()
			} else if operate == "addAtHead" {
				myLinkedList.AddAtHead(test.values[i][0])
			} else if operate == "addAtTail" {
				myLinkedList.AddAtTail(test.values[i][0])
			} else if operate == "addAtIndex" {
				myLinkedList.AddAtIndex(test.values[i][0], test.values[i][1])
			} else if operate == "get" {
				ret := myLinkedList.Get(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %v for Get(%d); expected %d", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "deleteAtIndex" {
				myLinkedList.DeleteAtIndex(test.values[i][0])
			}
		}
	}
}