package main

import (
	"leetcode/algorithms/1636.SortArraybyIncreasingFrequency/frequencySort"
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T)  {
	tests := []struct{
		nums []int
		output []int
	}{
		{[]int{1,1,2,2,2,3}, []int{3,1,1,2,2,2}},
		{[]int{2,3,1,3,2}, []int{1,3,3,2,2}},
		{[]int{-1,1,-6,4,5,-6,1,4,1}, []int{5,-1,4,4,-6,-6,1,1,1}},
	}

	for _, test := range tests {
		ret := frequencySort.FrequencySort(test.nums)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
		}
	}
}