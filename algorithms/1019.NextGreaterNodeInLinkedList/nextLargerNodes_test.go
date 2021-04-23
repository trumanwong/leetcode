package main

import (
	"leetcode/algorithms/1019.NextGreaterNodeInLinkedList/nextLargerNodes"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	tests := []struct {
		head   []int
		output []int
	}{
		{[]int{2, 1, 5}, []int{5, 5, 0}},
		{[]int{2, 7, 4, 3, 5}, []int{7, 0, 5, 5, 0}},
	}

	for _, test := range tests {
		ret := nextLargerNodes.NextLargerNodes(Constructor(test.head))
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}
