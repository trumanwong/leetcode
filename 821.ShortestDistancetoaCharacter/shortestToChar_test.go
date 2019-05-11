package main

import (
	"testing"
	"truman.com/leetcode/821.ShortestDistancetoaCharacter/shortestToChar"
)

func TestShortestToChar(t *testing.T)  {
	tests := []struct{
		S string
		c byte
		output []int
	}{
		{"loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
	}

	for _, test := range tests {
		ret := shortestToChar.ShortestToChar(test.S, test.c)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for S=%s, c=%c; expected %v", ret, test.S, test.c, test.output)
				break
			}
		}
	}
}