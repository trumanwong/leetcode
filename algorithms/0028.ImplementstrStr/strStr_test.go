package main

import (
	"leetcode/algorithms/0028.ImplementstrStr/strStr"
	"testing"
)

func TestStrStr(t *testing.T)  {
	tests := []struct{
		haystack string
		needle string
		output int
	} {
		{"hello","ll", 2},
		{"aaaaa", "bba", -1},
	}

	for _, test := range tests {
		ret := strStr.StrStr(test.haystack, test.needle)
		if ret != test.output {
			t.Errorf("Got %d for input %s, needle %s; expected %d", ret, test.haystack, test.needle, test.output)
		}
	}
}