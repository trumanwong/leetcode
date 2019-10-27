package main

import (
	"leetcode/algorithms/1160.FindWordsThatCanBeFormedbyCharacters/countCharacters"
	"testing"
)

func TestCountCharacters(t *testing.T)  {
	tests := []struct{
		words []string
		chars string
		output int
	}{
		{[]string{"cat","bt","hat","tree"}, "atach", 6},
		{[]string{"hello","world","leetcode"}, "welldonehoneyr", 10},
	}

	for _, test := range tests {
		ret := countCharacters.CountCharacters(test.words, test.chars)
		if ret != test.output {
			t.Errorf("Got %d for words = %v, chars = %s; expected %d", ret, test.words, test.chars, test.output)
		}
	}
}