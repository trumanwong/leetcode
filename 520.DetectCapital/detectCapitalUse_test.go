package main

import (
	"testing"
	"truman.com/leetcode/520.DetectCapital/detectCapitalUse"
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
		ret := detectCapitalUse.DetectCapitalUse(test.word)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.word, test.output)
		}
	}
}