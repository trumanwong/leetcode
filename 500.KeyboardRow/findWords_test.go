package main

import (
	"testing"
	"truman.com/leetcode/500.KeyboardRow/findWords"
)

func TestFindWords(t *testing.T)  {
	tests := []struct{
		words []string
		output []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
	}

	for _, test := range tests {
		ret := findWords.FindWords(test.words)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.words, test.output)
				break
			}
		}
	}
}