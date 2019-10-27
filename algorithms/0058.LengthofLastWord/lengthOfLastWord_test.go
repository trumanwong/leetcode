package main

import (
	"leetcode/algorithms/0058.LengthofLastWord/lengthOfLastWord"
	"testing"
)

func TestLengthOfLastWord(t *testing.T)  {
	tests := []struct{
		input string
		output int
	}{
		{"Hello World", 5},
	}

	for _, test := range tests {
		ret := lengthOfLastWord.LengthOfLastWord(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}