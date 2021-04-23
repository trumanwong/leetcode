package main

import (
	"leetcode/algorithms/0151.ReverseWordsinaString/reverseWords"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for _, test := range tests {
		ret := reverseWords.ReverseWords(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
