package main

import (
	"leetcode/algorithms/0792.NumberofMatchingSubsequences/numMatchingSubseq"
	"testing"
)

func TestNumMatchingSubseq(t *testing.T)  {
	tests := []struct{
		S string
		words []string
		output int
	}{
		{"abcde",[]string{"a", "bb", "acd", "ace"},3},
	}

	for _, test := range tests {
		ret := numMatchingSubseq.NumMatchingSubseq(test.S, test.words)
		if ret != test.output {
			t.Errorf("Got %d for input %s, words = %v; expected %d", ret, test.S, test.words, test.output)
		}
	}
}