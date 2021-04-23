package main

import (
	"leetcode/algorithms/0203.RemoveLinkedListElements/removeElements"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		head   []int
		val    int
		output []int
	}{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		ret := removeElements.RemoveElements(Constructor(test.head), test.val)
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v, val = %d", ret, test.head, test.val, test.output)
		}
	}
}
