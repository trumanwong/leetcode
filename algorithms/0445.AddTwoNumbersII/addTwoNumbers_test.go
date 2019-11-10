package main

import (
	"leetcode/algorithms/0445.AddTwoNumbersII/addTwoNumbers"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T)  {
	tests := []struct{
		l1 []int
		l2 []int
		output []int
	}{
		{[]int{7,2,4,3}, []int{5,6,4}, []int{7, 8, 0, 7}},
	}

	for _, test := range tests {
		ret := addTwoNumbers.AddTwoNumbers(Constructor(test.l1), Constructor(test.l2))
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for %v + %v; expected %v", ret, test.l1, test.l2, test.output)
		}
	}
}