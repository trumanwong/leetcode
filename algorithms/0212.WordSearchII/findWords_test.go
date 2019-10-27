package main

import (
	"leetcode/algorithms/0212.WordSearchII/findWords"
	"testing"
)

func TestFindWords(t *testing.T)  {
	tests := []struct{
		board [][]byte
		words []string
		output []string
	}{
		{[][]byte{
			{'o','a','a','n'},
			{'e','t','a','e'},
			{'i','h','k','r'},
			{'i','f','l','v'},
		}, []string{"oath","pea","eat","rain"}, []string{"oath", "eat"},},
	}

	for _, test := range tests {
		ret := findWords.FindWords(test.board, test.words)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for board = %v, words = %v; expected %v", ret, test.board, test.words, test.output)
				break
			}
		}
	}
}