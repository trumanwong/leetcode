package main

import (
	"leetcode/algorithms/1170.CompareStringsbyFrequencyoftheSmallestCharacter/numSmallerByFrequency"
	"reflect"
	"testing"
)

func TestNumSmallerByFrequency(t *testing.T)  {
	tests := []struct{
		queries []string
		words []string
		output []int
	}{
		{[]string{"cbd"}, []string{"zaaaz"}, []int{1}},
		{[]string{"bbb","cc"}, []string{"a","aa","aaa","aaaa"}, []int{1, 2}},
	}

	for _, test := range tests {
		ret := numSmallerByFrequency.NumSmallerByFrequency(test.queries, test.words)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for queries = %v, words = %v; expected %v", ret, test.queries, test.words, test.output)
		}
	}
}