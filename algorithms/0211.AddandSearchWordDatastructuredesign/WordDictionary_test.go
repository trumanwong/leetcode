package main

import (
	"leetcode/algorithms/0211.AddandSearchWordDatastructuredesign/WordDictionary"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	tests := []struct {
		operates []string
		values   [][]interface{}
		output   []interface{}
	}{
		{[]string{"WordDictionary","addWord","addWord","addWord","search","search","search","search"},
			[][]interface{}{{nil}, {"bad"},{"dad"},{"mad"},{"pad"},{"bad"},{".ad"},{"b.."}},
			[]interface{}{nil,nil,nil,nil,false,true,true,true},
		},
	}

	for _, test := range tests {
		wordDictionary := WordDictionary.WordDictionary{}
		for i, operate := range test.operates {
			if operate == "WordDictionary" {
				wordDictionary = WordDictionary.Constructor()
			} else if operate == "addWord" {
				wordDictionary.AddWord(test.values[i][0].(string))
			} else if operate == "search" {
				ret := wordDictionary.Search(test.values[i][0].(string))
				if ret != test.output[i] {
					t.Errorf("Got %t for Search(%s); expected %t", ret, test.values[i][0], test.output[i])
				}
			}
		}
	}
}
