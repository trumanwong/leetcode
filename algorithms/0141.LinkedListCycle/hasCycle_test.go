package main

import (
	"leetcode/algorithms/0141.LinkedListCycle/hasCycle"
	. "leetcode/common/list"
	"testing"
)

func TestHasCycle(t *testing.T)  {
	tests := []struct{
		head []int
		output bool
	}{
		{[]int{3, 2, 0, -4}, true},
		{[]int{1, 2}, true},
		{[]int{1}, false},
	}

	for _, test := range tests {
		ret := hasCycle.HasCycle(Constructor(test.head))
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.head, test.output)
		}
	}
}