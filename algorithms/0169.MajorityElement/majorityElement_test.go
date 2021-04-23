package main

import (
	"leetcode/algorithms/0169.MajorityElement/majorityElement"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		ret := majorityElement.MajorityElement(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
