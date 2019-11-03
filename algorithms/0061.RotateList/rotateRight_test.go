package main

import (
	"leetcode/algorithms/0061.RotateList/rotateRight"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T)  {
	tests := []struct{
		head []int
		k int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
	}

	for _, test := range tests {
		ret := rotateRight.RotateRight(Constructor(test.head), test.k)
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v, k = %d", ret, test.head, test.k, test.output)
		}
	}
}