package main

import (
	"leetcode/algorithms/0949.LargestTimeforGivenDigits/largestTimeFromDigits"
	"testing"
)

func TestLargestTimeFromDigits(t *testing.T) {
	tests := []struct {
		A      []int
		output string
	}{
		{[]int{1, 2, 3, 4}, "23:41"},
		{[]int{5, 5, 5, 5}, ""},
		{[]int{0, 0, 3, 0}, "03:00"},
	}

	for _, test := range tests {
		ret := largestTimeFromDigits.LargestTimeFromDigits(test.A)
		if ret != test.output {
			t.Errorf("Got %s for input %v; expected %s", ret, test.A, test.output)
		}
	}
}
