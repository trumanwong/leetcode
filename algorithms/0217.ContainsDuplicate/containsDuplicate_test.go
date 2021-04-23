package main

import (
	"leetcode/algorithms/0217.ContainsDuplicate/containsDuplicate"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, test := range tests {
		ret := containsDuplicate.ContainsDuplicate(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}
