package main

import (
	"leetcode/algorithms/0092.ReverseLinkedListII/reverseBetween"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head   []int
		m      int
		n      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
	}

	for _, test := range tests {
		ret := reverseBetween.ReverseBetween(Constructor(test.head), test.m, test.n)
		if !reflect.DeepEqual(test.output, ret.ToArray()) {
			t.Errorf("Got %v for input %v, m = %d, n = %d; expected %v", ret, test.head, test.m, test.n, test.output)
		}
	}
}
