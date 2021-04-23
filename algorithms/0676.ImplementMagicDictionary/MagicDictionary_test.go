package main

import (
	"leetcode/algorithms/0676.ImplementMagicDictionary/MagicDictionary"
	"testing"
)

func TestMagicDictionary(t *testing.T) {
	tests := []struct {
		operates []string
		values   [][]string
		output   []interface{}
	}{
		{[]string{"MagicDictionary", "buildDict", "search", "search", "search", "search"},
			[][]string{{}, {"hello", "leetcode"}, {"hello"}, {"hhllo"}, {"hell"}, {"leetcoded"}}, []interface{}{nil, nil, false, true, false, false}},
	}

	for _, test := range tests {
		dict := MagicDictionary.MagicDictionary{}
		for i, operate := range test.operates {
			if operate == "MagicDictionary" {
				dict = MagicDictionary.Constructor()
			} else if operate == "buildDict" {
				dict.BuildDict(test.values[i])
			} else if operate == "search" {
				ret := dict.Search(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for search(%s); expected %t", ret, test.values[i][0], test.output[i])
				}
			}
		}
	}
}
