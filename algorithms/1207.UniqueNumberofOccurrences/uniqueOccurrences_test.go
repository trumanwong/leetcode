package main

import (
	"leetcode/algorithms/1207.UniqueNumberofOccurrences/uniqueOccurrences"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr    []int
		output bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range tests {
		ret := uniqueOccurrences.UniqueOccurrences(test.arr)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.arr, test.output)
		}
	}
}
