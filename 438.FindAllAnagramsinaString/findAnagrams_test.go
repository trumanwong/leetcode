package main

import (
	"testing"
	"truman.com/leetcode/438.FindAllAnagramsinaString/findAnagrams"
)

func TestFindAnagrams(t *testing.T)  {
	tests := []struct{
		s string
		p string
		output []int
	}{
		{s: "cbaebabacd", p: "abc", output: []int{0, 6}},
		{s: "abab", p: "ab", output: []int{0, 1, 2}},
	}

	for _, test := range tests {
		ret := findAnagrams.FindAnagrams(test.s, test.p)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for s = %s, p = %s; expected %v", ret, test.s, test.p, test.output)
			}
		}
	}
}