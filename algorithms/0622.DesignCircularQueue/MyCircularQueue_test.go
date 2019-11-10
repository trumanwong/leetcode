package main

import (
	"leetcode/algorithms/0622.DesignCircularQueue/MyCircularQueue"
	"testing"
)

func TestMyCircularQueue(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output []interface{}
	}{
		{[]string{"MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue", "enQueue","Rear"},
			[][]int{{3},{1},{2},{3},{4},{},{},{},{4},{}},
			[]interface{}{nil,true,true,true,false,3,true,true,true,4}},
	}

	for _, test := range tests {
		queue := MyCircularQueue.MyCircularQueue{}
		for i, operate := range test.operates {
			if operate == "MyCircularQueue" {
				queue = MyCircularQueue.Constructor(test.values[i][0])
			} else if operate == "enQueue" {
				ret := queue.EnQueue(test.values[i][0])
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for enQueue(%d); expected %t", ret, test.values[i][0], test.output[i])
					break
				}
			} else if operate == "deQueue" {
				ret := queue.DeQueue()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for deQueue(); expected %t", ret, test.output[i])
					break
				}
			} else if operate == "Front" {
				ret := queue.Front()
				if ret != test.output[i].(int) {
					t.Errorf("Got %d for Front(); expected %d", ret, test.output[i])
					break
				}
			} else if operate == "Rear" {
				ret := queue.Rear()
				if ret != test.output[i].(int) {
					t.Errorf("Got %d for Rear(); expected %d", ret, test.output[i])
					break
				}
			} else if operate == "isEmpty" {
				ret := queue.IsEmpty()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for isEmpty(); expected %t", ret, test.output[i])
					break
				}
			} else if operate == "isFull" {
				ret := queue.IsFull()
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for isFull(); expected %t", ret, test.output[i])
					break
				}
			}
		}
	}
}