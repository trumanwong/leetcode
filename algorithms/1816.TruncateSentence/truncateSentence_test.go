package main

import (
	"leetcode/algorithms/1816.TruncateSentence/truncateSentence"
	"testing"
)

func TestTruncateSentence(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output string
	}{
		{"Hello how are you Contestant", 4, "Hello how are you"},
		{"What is the solution to this problem", 4, "What is the solution"},
		{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
	}

	for _, test := range tests {
		ret := truncateSentence.TruncateSentence(test.s, test.k)
		if ret != test.output {
			t.Errorf("Got %s for input (%s, %d); expected %s", ret, test.s, test.k, test.output)
		}
	}
}
