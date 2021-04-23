package main

import (
	"leetcode/algorithms/0208.ImplementTrie/Trie"
	"testing"
)

func TestTrie(t *testing.T) {
	tests := []struct {
		operates []string
		values   [][]string
		output   []interface{}
	}{
		{[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}, []interface{}{nil, nil, true, false, true, nil, true}},
	}

	for _, test := range tests {
		trie := Trie.Trie{}
		for i, operate := range test.operates {
			if operate == "Trie" {
				trie = Trie.Constructor()
			} else if operate == "insert" {
				trie.Insert(test.values[i][0])
			} else if operate == "search" {
				ret := trie.Search(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for search(%s); expected %t", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "startsWith" {
				ret := trie.StartsWith(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for startsWith(%s); expected %t", ret, test.values[i][0], test.output[i])
				}
			}
		}
	}
}
