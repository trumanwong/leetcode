package main

import (
	"leetcode/algorithms/1081.SmallestSubsequenceofDistinctCharacters/smallestSubsequence"
	"testing"
)

func TestSmallestSubsequence(t *testing.T) {
	tests := []struct {
		text   string
		output string
	}{
		{"cdadabcc", "adbc"},
		{"abcd", "abcd"},
		{"ecbacba", "eacb"},
		{"leetcode", "letcod"},
	}

	for _, test := range tests {
		ret := smallestSubsequence.SmallestSubsequence(test.text)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.text, test.output)
		}
	}
}
