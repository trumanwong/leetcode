package main

import (
	"leetcode/algorithms/0641.DesignCircularDeque/MyCircularDeque"
	"testing"
)

func TestMyCircularDeque(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output []interface{}
	}{
		{[]string{"MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"},
		[][]int{{3},{1},{2},{3},{4},{},{},{},{4},{}},
		[]interface{}{nil,true,true,true,false,2,true,true,true,4}},
	}

	for _, test := range tests {
		queue := MyCircularDeque.MyCircularDeque{}
		for i, operate := range test.operates {
			if operate == "MyCircularDeque" {
				queue = MyCircularDeque.Constructor(test.values[i][0])
			} else if operate == "insertFront" {
				ret := queue.InsertFront(test.values[i][0])
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for insertFront(%d); expected %t", ret, test.values[i][0], test.output[i])
					break
				}
			} else if operate == "insertLast" {
				ret := queue.InsertLast(test.values[i][0])
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for insertLast(%d); expected %t", ret, test.values[i][0], test.output[i])
					break
				}
			} else if operate == "deleteFront" {
				ret := queue.DeleteFront()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for DeleteFront(); expected %t", ret, test.output[i])
					break
				}
			} else if operate == "deleteLast" {
				ret := queue.DeleteLast()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for DeleteLast(); expected %t", ret, test.output[i])
					break
				}
			} else if operate == "getFront" {
				ret := queue.GetFront()
				if ret != test.output[i].(int) {
					t.Errorf("Got %d for getFront(); expected %d", ret, test.output[i])
					break
				}
			} else if operate == "getRear" {
				ret := queue.GetRear()
				if ret != test.output[i].(int) {
					t.Errorf("Got %d for getRear(); expected %d", ret, test.output[i])
					break
				}
			} else if operate == "isEmpty" {
				ret := queue.IsEmpty()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for isEmpty(); expected %t", ret, test.output[i])
					break
				}
			} else if operate == "IsFull" {
				ret := queue.IsFull()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for IsFull(); expected %t", ret, test.output[i])
					break
				}
			}
		}

	}
}