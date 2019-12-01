package main

import (
	"leetcode/algorithms/0147.InsertionSortList/insertionSortList"
	"leetcode/common/list"
	"reflect"
	"testing"
)

func TestInsertionSortList(t *testing.T)  {
	tests := []struct{
		head []int
		output []int
	}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		ret := insertionSortList.InsertionSortList(list.Constructor(test.head))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}