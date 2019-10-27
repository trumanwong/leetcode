package main

import (
	"leetcode/algorithms/0139.WordBreak/wordBreak"
	"testing"
)

func TestWordBreak(t *testing.T)  {
	tests := []struct{
		s string
		wordDict []string
		output bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, test := range tests {
		ret := wordBreak.WordBreak(test.s, test.wordDict)
		if ret != test.output {
			t.Errorf("Got %t for s = %s, wordDict = %v; expected %t", ret, test.s, test.wordDict, test.output)
		}
	}
}