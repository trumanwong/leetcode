package main

import (
	"leetcode/algorithms/0148.SortList/sortList"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T)  {
	tests := []struct{
		head []int
		output []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
	}

	for _, test := range tests {
		ret := sortList.SortList(Constructor(test.head))
		if !reflect.DeepEqual(test.output, ret.ToArray()) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}