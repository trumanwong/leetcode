package main

import (
	"leetcode/algorithms/0318.MaximumProductofWordLengths/maxProduct"
	"testing"
)

func TestMaxProduct(t *testing.T)  {
	tests := []struct{
		words []string
		output int
	}{
		{[]string{"abcw","baz","foo","bar","xtfn","abcdef"}, 16},
		{[]string{"a","ab","abc","d","cd","bcd","abcd"}, 4},
		{[]string{"a","aa","aaa","aaaa"},0},
	}

	for _, test := range tests {
		ret := maxProduct.MaxProduct(test.words)
		if ret != test.output {
			t.Errorf("Got %v for input %v; expected %d", ret, test.words, test.output)
		}
	}
}