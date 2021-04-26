package main

import (
	"leetcode/algorithms/5739.FrequencyoftheMostFrequentElement/maxFrequency"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{1, 2, 4}, 5, 3},
		{[]int{1, 4, 8, 13}, 5, 2},
	}
	for _, test := range tests {
		ret := maxFrequency.MaxFrequency(test.nums, test.k)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.nums, test.k, test.output)
		}
	}
}
