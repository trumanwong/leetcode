package main

import (
	"testing"
	"truman.com/leetcode/firstMissingPositive/firstMissingPositive"
)

func TestFirstMissingPositive(t *testing.T)  {
	tests := []struct {
		input []int
		output int
	}{
		{[]int{1,2,0}, 3},
		{[]int{3,4,-1,1}, 2},
		{[]int{7,8,9,11,12}, 1},
	}

	for _, test := range tests {
		ret := firstMissingPositive.FirstMissingPositive(test.input)

		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
