package main

import (
	"testing"
	detectCapitalUse2 "truman.com/leetcode/algorithms/520.DetectCapital/detectCapitalUse"
)

func TestDetectCapitalUse(t *testing.T)  {
	tests := []struct{
		word string
		output bool
	}{
		{"USA", true},
		{"FlaG", false},
	}

	for _, test := range tests {
		ret := detectCapitalUse2.DetectCapitalUse(test.word)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.word, test.output)
		}
	}
}