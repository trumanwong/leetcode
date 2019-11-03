package main

import (
	"leetcode/algorithms/0082.RemoveDuplicatesfromSortedListII/deleteDuplicates"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T)  {
	tests := []struct{
		head []int
		output []int
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{[]int{1, 1, 1, 2, 3}, []int{2, 3}},
	}

	for _, test := range tests {
		ret := deleteDuplicates.DeleteDuplicates(Constructor(test.head))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}