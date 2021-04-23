package main

import (
	"fmt"
	"leetcode/algorithms/0002.AddTwoNumbers/addTwoNumbers"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1     []int
		l2     []int
		output []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	}

	for _, test := range tests {
		ret := addTwoNumbers.AddTwoNumbers(Constructor(test.l1), Constructor(test.l2))
		fmt.Println(ret)
		if !reflect.DeepEqual(test.output, ret.ToArray()) {
			t.Errorf("Got %v for input %v + %v; expected %v", ret, test.l1, test.l2, test.output)
		}
	}
}
