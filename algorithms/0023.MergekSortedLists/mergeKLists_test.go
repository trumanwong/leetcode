package main

import (
	"leetcode/algorithms/0023.MergekSortedLists/mergeKLists"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists  [][]int
		output []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
	}

	for _, test := range tests {
		lists := []*ListNode{}
		for _, list := range test.lists {
			lists = append(lists, Constructor(list))
		}
		ret := mergeKLists.MergeKLists(lists)
		if !reflect.DeepEqual(test.output, ret.ToArray()) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.lists, test.output)
		}
	}
}
