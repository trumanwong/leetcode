package main

import (
	"fmt"
	"leetcode/algorithms/0049.GroupAnagrams/groupAnagrams"
	"testing"
)

func TestGroupAnagrams(t *testing.T)  {
	tests := []struct{
		strs []string
		output [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}},
	}

	for _, test := range tests {
		ret := groupAnagrams.GroupAnagrams(test.strs)
		fmt.Println(ret)
	}
}