package main

import (
	"leetcode/algorithms/0941.ValidMountainArray/validMountainArray"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		A      []int
		output bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
	}

	for _, test := range tests {
		ret := validMountainArray.ValidMountainArray(test.A)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.A, test.output)
		}
	}
}
