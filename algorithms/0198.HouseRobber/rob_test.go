package main

import (
	"leetcode/algorithms/0198.HouseRobber/rob"
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{1, 8, 1, 1, 6}, 14},
	}

	for _, test := range tests {
		ret := rob.Rob(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
