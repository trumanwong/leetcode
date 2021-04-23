package main

import (
	"leetcode/algorithms/0692.TopKFrequentWords/topKFrequent"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		words  []string
		k      int
		output []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2,
			[]string{"i", "love"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4,
			[]string{"the", "is", "sunny", "day"}},
	}

	for _, test := range tests {
		ret := topKFrequent.TopKFrequent(test.words, test.k)
		if len(ret) != len(test.output) {
			t.Errorf("Got %v for input %v, k = %d; expected %v", ret, test.words, test.k, test.output)
			continue
		}
		for _, v := range ret {
			if !inArray(v, test.output) {
				t.Errorf("Got %v for input %v, k = %d; expected %v", ret, test.words, test.k, test.output)
				break
			}
		}
	}
}

func inArray(needle string, arr []string) bool {
	for _, v := range arr {
		if v == needle {
			return true
		}
	}
	return false
}
