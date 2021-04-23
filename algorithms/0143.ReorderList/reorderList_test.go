package main

import (
	"leetcode/algorithms/0143.ReorderList/reorderList"
	"leetcode/common/list"
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		head   []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	}

	for _, test := range tests {
		head := list.Constructor(test.head)
		reorderList.ReorderList(head)
		if !reflect.DeepEqual(head.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", head, test.head, test.output)
		}
	}
}
