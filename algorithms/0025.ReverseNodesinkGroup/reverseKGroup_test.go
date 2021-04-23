package main

import (
	"leetcode/algorithms/0025.ReverseNodesinkGroup/reverseKGroup"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head   []int
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	for _, test := range tests {
		ret := reverseKGroup.ReverseKGroup(Constructor(test.head), test.k)
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v, k = %d; expected %v", ret, test.head, test.k, test.output)
		}
	}
}
