package main

import (
	"leetcode/algorithms/0748.ShortestCompletingWord/shortestCompletingWord"
	"testing"
)

func TestShortestCompletingWord(t *testing.T) {
	tests := []struct {
		licensePlate string
		words        []string
		output       string
	}{
		{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps"},
		{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps"},
	}

	for _, test := range tests {
		ret := shortestCompletingWord.ShortestCompletingWord(test.licensePlate, test.words)
		if ret != test.output {
			t.Errorf("Got %s for input licensePlate = %s, words = %v; expected %s", ret, test.licensePlate, test.words, test.output)
		}
	}
}
