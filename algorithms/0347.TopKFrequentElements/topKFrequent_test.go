package main

import (
	"leetcode/algorithms/0347.TopKFrequentElements/topKFrequent"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		ret := topKFrequent.TopKFrequent(test.nums, test.k)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v, k = %d; expected %v", ret, test.nums, test.k, test.output)
				break
			}
		}
	}
}
