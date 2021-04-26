package main

import (
	"leetcode/interview_questions/BinarySearch/04search/search"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{3,1}, 1, 1},
	}

	for _, test := range tests {
		ret := search.Search(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.nums, test.target, test.output)
		}
	}
}
