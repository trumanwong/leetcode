package main

import (
	"leetcode/algorithms/0024.SwapNodesinPairs/swapPairs"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T)  {
	tests := []struct{
		head []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
	}

	for _, test := range tests {
		ret := swapPairs.SwapPairs(Constructor(test.head))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.head, test.output)
		}
	}
}