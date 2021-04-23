package main

import (
	"leetcode/algorithms/0520.DetectCapital/detectCapitalUse"
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		word   string
		output bool
	}{
		{"USA", true},
		{"FlaG", false},
	}

	for _, test := range tests {
		ret := detectCapitalUse.DetectCapitalUse(test.word)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.word, test.output)
		}
	}
}
