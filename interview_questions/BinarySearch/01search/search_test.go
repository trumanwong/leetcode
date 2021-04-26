package main

import (
	"leetcode/interview_questions/BinarySearch/01search/search"
	"testing"
)

func TestSearch(t *testing.T)  {
	tests := []struct{
		nums []int
		target int
		output int
	}{
		{[]int{-1,0,3,5,9,12}, 9, 4},
		{[]int{-1,0,3,5,9,12}, 2, -1},
	}

	for _, test := range tests {
		ret := search.Search(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.nums, test.target, test.output)
		}
	}
}