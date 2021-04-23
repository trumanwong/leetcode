package main

import (
	"leetcode/algorithms/0206.ReverseLinkedList/reverseList"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		head   []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		ret := reverseList.ReverseList(Constructor(test.head))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}
