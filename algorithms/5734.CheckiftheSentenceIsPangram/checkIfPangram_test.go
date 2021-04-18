package main

import (
	"leetcode/algorithms/5734.CheckiftheSentenceIsPangram/checkIfPangram"
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	tests := []struct {
		sentence string
		output   bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
	}

	for _, test := range tests {
		ret := checkIfPangram.CheckIfPangram(test.sentence)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.sentence, test.output)
		}
	}
}
