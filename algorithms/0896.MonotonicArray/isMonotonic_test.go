package main

import (
	"leetcode/algorithms/0896.MonotonicArray/isMonotonic"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		A      []int
		output bool
	}{
		{[]int{1, 2, 2, 3}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 2, 4, 5}, true},
		{[]int{1, 1, 1}, true},
	}

	for _, test := range tests {
		ret := isMonotonic.IsMonotonic(test.A)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.A, test.output)
		}
	}
}
