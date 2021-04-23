package main

import (
	"leetcode/algorithms/0328.OddEvenLinkedList/oddEvenList"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		head   []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	}

	for _, test := range tests {
		ret := oddEvenList.OddEvenList(Constructor(test.head))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}
