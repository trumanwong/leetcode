package main

import (
	"leetcode/algorithms/0019.RemoveNthNodeFromEndofList/removeNthFromEnd"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head   []int
		n      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
	}

	for _, test := range tests {
		head := Constructor(test.head)
		removeNthFromEnd.RemoveNthFromEnd(head, test.n)
		res := head.ToArray()
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("Got %v for input %v, n = %d; expected %v", res, test.head, test.n, test.output)
		}
	}
}
