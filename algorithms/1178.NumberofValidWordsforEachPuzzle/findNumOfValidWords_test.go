package main

import (
	"leetcode/algorithms/1178.NumberofValidWordsforEachPuzzle/findNumOfValidWords"
	"testing"
)

func TestFindNumOfValidWords(t *testing.T)  {
	tests := []struct{
		words []string
		puzzles []string
		output []int
	}{
		{[]string{"aaaa","asas","able","ability","actt","actor","access"},
			[]string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"},
			[]int{1,1,3,2,4,0}},
	}

	for _, test := range tests {
		ret := findNumOfValidWords.FindNumOfValidWords(test.words, test.puzzles)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for words = %v, puzzles = %v; expected %v", ret, test.words, test.puzzles, test.output)
				break
			}
		}
	}
}