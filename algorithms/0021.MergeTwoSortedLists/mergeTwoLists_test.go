package main

import (
	"leetcode/algorithms/0021.MergeTwoSortedLists/mergeTwoLists"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     []int
		l2     []int
		output []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	}

	for _, test := range tests {
		ret := mergeTwoLists.MergeTwoLists(Constructor(test.l1), Constructor(test.l2))
		if !reflect.DeepEqual(test.output, ret.ToArray()) {
			t.Errorf("Got %v for merge(%v, %v); expected %v", ret, test.l1, test.l2, test.output)
		}
	}
}
