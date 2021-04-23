package main

import (
	"leetcode/algorithms/0744.FindSmallestLetterGreaterThanTarget/nextGreatestLetter"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		output  byte
	}{
		{[]byte("cfj"), 'a', 'c'},
		{[]byte("cfj"), 'c', 'f'},
		{[]byte("cfj"), 'd', 'f'},
		{[]byte("cfj"), 'g', 'j'},
		{[]byte("cfj"), 'j', 'c'},
		{[]byte("cfj"), 'k', 'c'},
	}

	for _, test := range tests {
		ret := nextGreatestLetter.NextGreatestLetter(test.letters, test.target)
		if ret != test.output {
			t.Errorf("Got %c for input %s, target %c; expected %c", ret, test.letters, test.target, test.output)
		}
	}
}
